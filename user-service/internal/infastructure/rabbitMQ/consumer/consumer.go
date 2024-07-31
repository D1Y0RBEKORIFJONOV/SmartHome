package consumer

import (
	"context"
	"encoding/json"
	user1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/user"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"user_service_smart_home/internal/config"
	"user_service_smart_home/internal/entity"
	clientgrpcserver "user_service_smart_home/internal/infastructure/client_grpc_server"
)

type Consumer struct {
	channel      *amqp091.Channel
	clientDevice clientgrpcserver.ServiceClient
}

func NewConsumer(cfg *config.Config) (*Consumer, error) {
	conn, err := amqp091.Dial(cfg.RabbitMQURL)
	if err != nil {
		return nil, err
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
		"user",
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
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			err := json.Unmarshal(d.Body, &ReqType)
			if err != nil {
				continue
			}

			switch ReqType.MethodName {
			case "user.create":
				consumer.createUser(d.Body)
			case "user.update":
				consumer.updateUser(d.Body)
			case "user.delete":
				consumer.deleteUser(d.Body)
			case "user.update_email":
				consumer.updateEmail(d.Body)
			case "user.update_password":
				consumer.updatePassword(d.Body)
			case "user.delete_user":
				consumer.deleteUser(d.Body)
			}
		}
	}()
	<-forever
	return nil
}

func (consumer *Consumer) createUser(body []byte) {
	var createUserReq entity.CreateUserReq
	if err := json.Unmarshal(body, &createUserReq); err != nil {
		log.Printf("err:%s", err.Error())
		return
	}
	_, err := consumer.clientDevice.UserServiceClient().CreateUser(context.Background(), &user1.CreateUSerReq{
		FirstName: createUserReq.FirstName,
		LastName:  createUserReq.LastName,
		Email:     createUserReq.Email,
		Password:  createUserReq.Password,
		Address:   createUserReq.Address,
	})
	if err != nil {
		log.Printf("err:%s", err.Error())
	}
}

func (consumer *Consumer) updateUser(body []byte) {
	var updateUserReq entity.UpdateUserReq
	if err := json.Unmarshal(body, &updateUserReq); err != nil {
		log.Printf("err:%s", err.Error())
		return
	}
	_, err := consumer.clientDevice.UserServiceClient().UpdateUser(context.Background(), &user1.UpdateUserReq{
		UserId:    updateUserReq.UserID,
		FirstName: updateUserReq.FirstName,
		LastName:  updateUserReq.LastName,
	})
	if err != nil {
		log.Printf("err:%s", err.Error())
	}
}
func (consumer *Consumer) deleteUser(body []byte) {
	var deleteUserReq entity.DeleteUserReq
	if err := json.Unmarshal(body, &deleteUserReq); err != nil {
		log.Printf("err:%s", err.Error())
	}
	_, err := consumer.clientDevice.UserServiceClient().DeleteUser(context.Background(), &user1.DeleteUserReq{
		UserId:       deleteUserReq.UserID,
		IsHardDelete: deleteUserReq.IsHardDelete,
	})
	if err != nil {
		log.Printf("err:%s", err.Error())
	}
}

func (consumer *Consumer) updateEmail(body []byte) {
	var updateEmailReq entity.UpdateEmailReq
	if err := json.Unmarshal(body, &updateEmailReq); err != nil {
		log.Printf("err:%s", err.Error())
	}
	_, err := consumer.clientDevice.UserServiceClient().UpdateEmail(context.Background(), &user1.UpdateEmailReq{
		UserId:   updateEmailReq.UserID,
		NewEmail: updateEmailReq.NewEmail,
	})
	if err != nil {
		log.Printf("err:%s", err.Error())
	}
}
func (consumer *Consumer) updatePassword(body []byte) {
	var updatePasswordReq entity.UpdatePasswordReq
	if err := json.Unmarshal(body, &updatePasswordReq); err != nil {
		log.Printf("err:%s", err.Error())
	}
	_, err := consumer.clientDevice.UserServiceClient().UpdatePassword(context.Background(), &user1.UpdatePasswordReq{
		UserId:      updatePasswordReq.UserID,
		NewPassword: updatePasswordReq.NewPassword,
		Password:    updatePasswordReq.Password,
	})
	if err != nil {
		log.Printf("err:%s", err.Error())
	}

}
