package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/tamiresviegas/warehouse/configs"

	"github.com/go-chi/chi"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
	"github.com/tamiresviegas/warehouse/handlers"
)

func main() {

	messageBroker()

	fmt.Println("Entrei aqui")

	err := configs.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.GetAll)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}

func messageBroker() {
	// PRODUCER:
	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "test",
	}

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("mensagem"),
		Headers: []protocol.Header{
			{
				Key:   "session",
				Value: []byte("123"),
			},
		},
	})

	if err != nil {
		log.Fatal("cannot write a message: ", err)
	}

	// CONSUMER:
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer",
		Topic:    "test",
		MinBytes: 0,
		MaxBytes: 10e6, // 10MB
	})

	for i := 0; i < 1; i++ {
		message, err := reader.ReadMessage(context.Background())

		for _, val := range message.Headers {
			if val.Key == "session" && string(val.Value) == "123" {
				fmt.Print("sessao correta")
			}
		}

		if err != nil {
			log.Fatal("cannot receive a message: ", err)
			reader.Close()
		}

		fmt.Print("receive a message: ", string(message.Value))
	}

	reader.Close()
}
