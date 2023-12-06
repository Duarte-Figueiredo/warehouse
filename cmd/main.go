package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tamiresviegas/warehouse/configs"

	"github.com/go-chi/chi"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
	"github.com/tamiresviegas/warehouse/handlers"
)

func main() {

	fmt.Println("Started warehouse app")

	err := configs.Load()
	if err != nil {
		log.Fatal(err)
	}

	kafkaUrl := GetEnvDefault("KAFKA_URL", "")    // kafka:9092
	topicName := GetEnvDefault("KAFKA_TOPIC", "") //test

	subscribeToKafka(kafkaUrl, topicName)

	startApi()
}

func startApi() {
	r := chi.NewRouter()
	r.Get("/products/", handlers.GetAll)                                           // "Clients should be able to see a list of available products in the warehouse."
	r.Get("/products/{category}/{brand}/{maxPrice}", handlers.GetProductsFiltered) // "Clients should be able to get products based on product category, brand and maximum price"
	r.Put("/products/{orderPId}", handlers.GetOrderP)                              // "Clients should be able to receive products from the warehouse."
	r.Get("/orders/", handlers.GetAllOrders)                                       // "Clients should be able to see the received orders"
	r.Post("/orders/", handlers.AddOrder)                                          // "Clients should be able to add products to the data base"
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}

func subscribeToKafka(kafkaUrl string, topicName string) {

	if kafkaUrl == "" {
		fmt.Println("Skipping kafka initialization due to missing 'KAFKA_URL' environment variable")
		return
	}

	if topicName == "" {
		fmt.Println("Skipping kafka initialization due to missing 'KAFKA_TOPIC' environment variable")
		return
	}

	// wait for kafka to be up and running
	//TODO update docker compose to wait for topic to be created
	time.Sleep(30 * time.Second)

	fmt.Println("connecting to " + kafkaUrl)

	// PRODUCER:
	writer := &kafka.Writer{
		Addr:  kafka.TCP(kafkaUrl),
		Topic: topicName,
	}

	fmt.Println("Writing test message to topic " + topicName)
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
		Brokers:  []string{kafkaUrl},
		GroupID:  "consumer",
		Topic:    topicName,
		MinBytes: 0,
		MaxBytes: 10e6, // 10MB
	})

	for i := 0; i < 1; i++ {
		message, err := reader.ReadMessage(context.Background())

		for _, val := range message.Headers {
			if val.Key == "session" && string(val.Value) == "123" {
				fmt.Println("correct session")
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

func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)

	if !ex {
		return defVal
	}

	return val
}
