package main

import (
	"log"
	"user_service_smart_home/internal/config"
	"user_service_smart_home/internal/infastructure/rabbitMQ/consumer"
)

func main() {
	cfg := config.New()
	c, err := consumer.NewConsumer(cfg)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Connecting to RabbitMQ")
	log.Println("Starting consumer")
	log.Fatal(c.Consume())
}
