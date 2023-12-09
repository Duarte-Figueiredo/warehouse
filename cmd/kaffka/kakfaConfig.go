package kaffka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartKafka(kafkaUrl string, topicName string) {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaUrl},
		GroupID:  "consumer",
		Topic:    topicName,
		MinBytes: 0,
		MaxBytes: 10e6, // 10MB
	})

	for {
		message, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Fatal("cannot receive a message: ", err)
			reader.Close()
		}

		// Quando a mensagem é recebida tenho de chamar o handler que dá update à tabela dos produtos
		fmt.Println("receive a message: ", string(message.Value)) // [{"product_id":"product1","quantity":3},{"product_id":"product2","quantity":2},{"product_id":"product3","quantity":1}]
	}

}
