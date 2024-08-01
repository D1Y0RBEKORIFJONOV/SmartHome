package consumer

import (
	"context"
	"device_service/internal/config"
	alarm_entity "device_service/internal/entity/alarm"
	speaker_entity "device_service/internal/entity/speaker"
	tv_entity "device_service/internal/entity/tv"
	clientgrpcserver "device_service/internal/infastructure/client_grpc_server"
	"encoding/json"
	tv1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/device/TV"
	alarm1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/device/smart_alarm"
	speaker1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/device/speaker"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Consumer struct {
	channel      *amqp091.Channel
	clientDevice clientgrpcserver.ServiceClient
}

func NewConsumer(cfg *config.Config) (*Consumer, error) {
	var err error
	var conn *amqp091.Connection
	for i := 0; i < 10; i++ {
		conn, err = amqp091.Dial(cfg.RabbitMQURL)
		if err != nil {
			log.Println("Failed to connect to RabbitMQ")
			time.Sleep(1 * time.Millisecond)
			continue
		}
	}
	client, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	clientDevice, err := clientgrpcserver.NewService(cfg)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		channel:      client,
		clientDevice: clientDevice,
	}, nil
}

func (consumer *Consumer) Consume() error {
	err := consumer.channel.ExchangeDeclare(
		"logs",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	q, err := consumer.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = consumer.channel.QueueBind(
		q.Name,
		"devices",
		"logs",
		false,
		nil,
	)
	if err != nil {
		return err
	}

	var ReqType struct {
		MethodName string `json:"method_name"`
	}

	msgs, err := consumer.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			err := json.Unmarshal(d.Body, &ReqType)
			if err != nil {
				continue
			}

			switch ReqType.MethodName {
			case "TV.AddTvToUser":
				consumer.handleAddTvToUser(d.Body)
			case "TV.AddChannel":
				consumer.handleAddChannel(d.Body)
			case "TV.DeleteChannel":
				consumer.handleDeleteChannel(d.Body)
			case "SPEAKER.AddSpeaker":
				consumer.handleAddSpeaker(d.Body)
			case "SPEAKER.AddChannel":
				consumer.handleAddSpeakerChannel(d.Body)
			case "SPEAKER.DeleteChannel":
				consumer.handleDeleteSpeakerChannel(d.Body)
			case "Alarm.AddSmartAlarm":
				consumer.handleAddSmartAlarm(d.Body)
			case "Alarm.CreateAlarmClock":
				consumer.handleCreateAlarmClock(d.Body)
			default:
			}
		}
	}()

	select {}
}

func (consumer *Consumer) handleAddTvToUser(body []byte) {
	var req tv_entity.AddTVReq
	log.Printf("%v", string(body))
	if err := json.Unmarshal(body, &req); err != nil {
		return
	}
	log.Printf("\n%s\n", req)
	_, err := consumer.clientDevice.TvService().AddTV(context.Background(), &tv1.AddTVReq{
		UserId:    req.UserID,
		ModelName: req.ModelName,
	})
	if err != nil {
		log.Printf("err: %v", err.Error())
	}
}

func (consumer *Consumer) handleAddChannel(body []byte) {
	var req tv_entity.AddChannelReq
	if err := json.Unmarshal(body, &req); err != nil {
		return
	}

	_, err := consumer.clientDevice.TvService().AddChannel(context.Background(), &tv1.AddChannelReq{
		UserId:      req.UserID,
		ChannelName: req.ChannelName,
	})
	if err != nil {
		log.Printf("err: %v", err.Error())
	}
}

func (consumer *Consumer) handleDeleteChannel(body []byte) {
	var req tv_entity.DeleteChannelReq
	if err := json.Unmarshal(body, &req); err != nil {
		return
	}

	_, err := consumer.clientDevice.TvService().DeleteChannel(context.Background(), &tv1.DeleteChannelReq{
		UserId:      req.UserID,
		ChannelName: req.ChannelName,
	})
	if err != nil {
		log.Printf("err: %v", err.Error())
	}
}

func (consumer *Consumer) handleAddSpeaker(body []byte) {
	var req speaker_entity.AddSpeakerReq
	if err := json.Unmarshal(body, &req); err != nil {
		return
	}

	_, err := consumer.clientDevice.SpeakerService().AddSpeaker(context.Background(), &speaker1.AddSpeakerReq{
		UserId:    req.UserID,
		ModelName: req.ModelName,
	})
	if err != nil {
		log.Printf("err: %v", err.Error())
	}
}

func (consumer *Consumer) handleAddSpeakerChannel(body []byte) {
	var req speaker_entity.AddChannelReq
	if err := json.Unmarshal(body, &req); err != nil {
		return
	}

	_, err := consumer.clientDevice.SpeakerService().AddChannel(context.Background(), &speaker1.AddChannelReqS{
		UserId:      req.UserID,
		ChannelName: req.ChannelName,
	})
	if err != nil {
		log.Printf("err: %v", err.Error())
	}
}

func (consumer *Consumer) handleDeleteSpeakerChannel(body []byte) {
	var req speaker_entity.DeleteChannelReq
	if err := json.Unmarshal(body, &req); err != nil {
		return
	}

	_, err := consumer.clientDevice.SpeakerService().DeleteChannel(context.Background(), &speaker1.DeleteChannelReqS{
		UserId:      req.UserID,
		ChannelName: req.ChannelName,
	})
	if err != nil {
		log.Printf("err: %v", err.Error())
	}
}

func (consumer *Consumer) handleAddSmartAlarm(body []byte) {
	var req alarm_entity.AddSmartAlarmReq
	if err := json.Unmarshal(body, &req); err != nil {
		return
	}

	_, err := consumer.clientDevice.SmartAlarm().AddSmartAlarm(context.Background(), &alarm1.AddSmartAlarmReq{
		UserId:     req.UserID,
		ModelName:  req.ModelName,
		DeviceName: req.DeviceName,
	})
	if err != nil {
		log.Printf("err: %v", err.Error())
	}
}

func (consumer *Consumer) handleCreateAlarmClock(body []byte) {
	var req alarm_entity.CreateAlarmClockReq
	if err := json.Unmarshal(body, &req); err != nil {
		return
	}

	_, err := consumer.clientDevice.SmartAlarm().CreateAlarmClock(context.Background(), &alarm1.CreateAlarmClockReq{
		UserId:     req.UserID,
		DeviceName: req.DeviceName,
		ClockTime:  req.ClockTime,
	})
	if err != nil {
		log.Printf("err: %v", err.Error())
	}
}
