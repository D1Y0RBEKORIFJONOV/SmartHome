package mongo_speaker

import (
	"context"
	"device_service/internal/config"
	err_entity "device_service/internal/entity/errors"
	speaker_entity "device_service/internal/entity/speaker"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	mongoClient *mongo.Client
	db          *mongo.Database
	collection  *mongo.Collection
}

func NewMongoDB(cfg *config.Config) (*MongoDB, error) {
	uri := "mongodb://" + cfg.DB.Host + cfg.DB.Port
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return &MongoDB{
		mongoClient: client,
		db:          client.Database(cfg.DB.Name),
		collection:  client.Database(cfg.DB.Name).Collection(cfg.DB.SpeakerCollection),
	}, nil
}

func (m *MongoDB) SaveUsersSpeaker(ctx context.Context, req *speaker_entity.Speaker) error {
	_, err := m.collection.InsertOne(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) SaveChannelsSpeaker(ctx context.Context, req *speaker_entity.Channel, userID string) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$push": bson.M{"channels": req},
	}

	_, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) GetUsersChannel(ctx context.Context, req *speaker_entity.GetUserChannelReq) (*speaker_entity.GetUserChannelRes, error) {
	filter := bson.M{"user_id": req.UserID}
	projection := bson.M{"channels": 1, "_id": 0}
	result, err := m.collection.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	var channels []speaker_entity.Channel
	for result.Next(ctx) {
		var userChannels struct {
			Channels []speaker_entity.Channel `bson:"channels"`
		}
		err := result.Decode(&userChannels)
		if err != nil {
			return nil, err
		}
		channels = append(channels, userChannels.Channels...)
	}
	if err := result.Err(); err != nil {
		return nil, err
	}
	return &speaker_entity.GetUserChannelRes{
		Channels: channels,
		Count:    int64(len(channels)),
	}, nil
}

func (m *MongoDB) GetCursorChannel(ctx context.Context, userID string) (*int, error) {
	filter := bson.M{"user_id": userID}
	projection := bson.M{"cursor_channel": 1, "_id": 0}

	var result struct {
		CursorChannel uint8 `bson:"cursor_channel"`
	}
	err := m.collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err_entity.ErrorNotFound
		}
		return nil, err
	}

	cursorChannelInt := int(result.CursorChannel)
	return &cursorChannelInt, nil
}

func (s *MongoDB) GetSoundSpeaker(ctx context.Context, userID string) (*int, error) {
	filter := bson.M{"user_id": userID}
	projection := bson.M{"sound": 1, "_id": 0}
	var result struct {
		Sound int `bson:"sound"`
	}

	err := s.collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err_entity.ErrorNotFound
		}
		return nil, err
	}
	return &result.Sound, nil
}

func (s *MongoDB) IsOnUsersSpeaker(ctx context.Context, userID string) (bool, error) {
	filter := bson.M{"user_id": userID}
	projection := bson.M{"on": 1, "_id": 0}
	var result struct {
		On bool `bson:"on"`
	}
	err := s.collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return result.On, nil
}

func (s *MongoDB) UpdateSound(ctx context.Context, userID string, val uint8) (*int, error) {
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$set": bson.M{
			"sound": val,
		},
	}
	_, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err_entity.ErrorNotFound
		}
		return nil, err
	}

	var res = int(val)
	return &res, nil
}

func (m *MongoDB) UpdateCursor(ctx context.Context, userID string, val uint8) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$set": bson.M{
			"cursor_channel": val,
		},
	}
	_, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return err_entity.ErrorNotFound
		}
		return nil
	}

	return nil
}

func (m *MongoDB) UpdateSpeakerOn(ctx context.Context, userID string, val bool) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$set": bson.M{
			"on": val,
		},
	}
	_, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return err_entity.ErrorNotFound
		}
	}
	return nil
}

func (m *MongoDB) DeleteChannel(ctx context.Context, userID, channelName string) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{"$pull": bson.M{"channels": bson.M{"channel_name": channelName}}}

	result, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return err_entity.ErrorNotFound
	}

	return nil
}
