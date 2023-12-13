package kaffka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/tamiresviegas/warehouse/handlers"
)

func StartKafkaReader(kafkaUrl string, topicSold string, topicSend string) {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaUrl},
		Topic:    topicSold,
		MinBytes: 0,
		MaxBytes: 10e6, // 10MB
	})

	for {
		message, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Fatal("cannot receive a message: ", err)
			reader.Close()
		}

		handlers.UpdateProducts(kafkaUrl, message, topicSend)

	}

}
