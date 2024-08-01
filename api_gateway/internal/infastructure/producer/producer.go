package producer

import (
	"api_gate_way/internal/config"
	"context"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Producer struct {
	channel *amqp091.Channel
}

func NewProducer(cfg config.Config) (*Producer, error) {
	producer := Producer{}
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		conn, err := amqp091.Dial(cfg.RabbitMQURL)
		if err != nil {
			log.Printf("Failed to open a channel: %s", err)
			continue
		}
		ch, err := conn.Channel()
		if err != nil {
			log.Printf("Failed to open a channel: %s", err)
		}

		RegisterProducer(ch)
		producer.channel = ch
	}

	return &producer, nil
}

func RegisterProducer(channel *amqp091.Channel) *Producer {
	return &Producer{channel: channel}
}

func (producer *Producer) Pub(ctx context.Context, req []byte, key string) error {
	err := producer.channel.ExchangeDeclare(
		"logs",
		"direct",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return err
	}
	err = producer.channel.PublishWithContext(ctx, "logs",
		key, false, false, amqp091.Publishing{
			ContentType: "application/json",
			Body:        req,
		})
	if err != nil {
		return err
	}
	return nil
}
