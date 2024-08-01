package main

import (
	"device_service/internal/config"
	"device_service/internal/infastructure/rammbitMQ/consumer"
	"log"
)

func main() {
	cfg := config.New()
	consumers, err := consumer.NewConsumer(cfg)
	if err != nil {
		log.Println(err)
	}
	log.Println("consumer started")
	log.Fatal(consumers.Consume())
}
