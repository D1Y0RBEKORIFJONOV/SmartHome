package mongo_alarmongo_alarm

import (
	"context"
	"device_service/internal/config"
	alarm_entity "device_service/internal/entity/alarm"
	err_entity "device_service/internal/entity/errors"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
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
		collection:  client.Database(cfg.DB.Name).Collection(cfg.DB.AlarmCollectionName),
	}, nil
}

func (m *MongoDB) SaveAlarm(ctx context.Context, alarm alarm_entity.SmartAlarm) error {
	_, err := m.collection.InsertOne(ctx, alarm)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) SaveAlarmClock(ctx context.Context, req *alarm_entity.Alarm, userId, deviceName string) error {
	filter := bson.M{"device_name": deviceName, "user_id": userId}
	update := bson.M{
		"$push": bson.M{
			"alarms": bson.M{
				"alarm_time":     req.AlarmTime.Format(time.RFC3339),
				"remaining_time": req.RemainingTime.Format(time.RFC3339),
			},
		},
	}

	fmt.Printf("AlarmTime: %v, RemainingTime: %v\n", req.AlarmTime, req.RemainingTime)

	result, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println("Error during update:", err)
		return err
	}

	if result.MatchedCount == 0 {
		fmt.Println("No matching documents found.")
		return err_entity.ErrorNotFound
	}

	fmt.Println("Update successful. ModifiedCount:", result.ModifiedCount)
	return nil
}

func (m *MongoDB) UpdateCurtain(ctx context.Context, userId, deviceName string, val bool) error {
	filter := bson.M{"user_id": userId, "device_name": deviceName}
	update := bson.M{
		"$set": bson.M{
			"curtain": val,
		},
	}
	re, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if re.MatchedCount == 0 {
		return err_entity.ErrorNotFound
	}
	return nil
}

func (m *MongoDB) UpdateDoor(ctx context.Context, userId, deviceName string, val bool) error {
	filter := bson.M{"user_id": userId, "device_name": deviceName}
	update := bson.M{
		"$set": bson.M{
			"door": val,
		},
	}
	r, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {

		return err
	}
	if r.MatchedCount == 0 {
		return err_entity.ErrorNotFound
	}
	return nil
}

func (m *MongoDB) GetAlarmUser(ctx context.Context, req *alarm_entity.RemainingTimeReq) (alarm_entity.RemainingTimRes, error) {
	filter := bson.M{"device_name": req.DeviceName, "user_id": req.UserID}
	projection := bson.M{"alarms": 1, "_id": 0}
	fmt.Println(filter)

	res, err := m.collection.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return alarm_entity.RemainingTimRes{}, err_entity.ErrorNotFound
		}
		return alarm_entity.RemainingTimRes{}, err
	}

	var result alarm_entity.RemainingTimRes
	for res.Next(ctx) {
		var doc struct {
			Alarms []alarm_entity.Alarm `bson:"alarms"`
		}
		err = res.Decode(&doc)
		if err != nil {
			return alarm_entity.RemainingTimRes{}, err
		}
		result.Alarms = append(result.Alarms, doc.Alarms...)
	}

	if err := res.Err(); err != nil {
		return alarm_entity.RemainingTimRes{}, err
	}

	result.Count = int64(len(result.Alarms))
	return result, nil
}

func (m *MongoDB) IsDeviceExists(ctx context.Context, userId, deviceName string) (bool, error) {
	fmt.Println(userId, deviceName)
	filter := bson.M{"user_id": userId, "device_name": deviceName}

	count, err := m.collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
