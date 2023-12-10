package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tamiresviegas/warehouse/cmd/kaffka"
	"github.com/tamiresviegas/warehouse/handlers"

	"github.com/tamiresviegas/warehouse/configs"

	"github.com/go-chi/chi"
)

func main() {

	fmt.Println("Started warehouse app")

	err := configs.Load()
	if err != nil {
		log.Fatal(err)
	}

	kafkaUrl := GetEnvDefault("KAFKA_URL", "") // kafka:9092
	topicSend := GetEnvDefault("KAFKA_TOPIC_SEND", "")
	topicSold := GetEnvDefault("KAFKA_TOPIC_SOLD", "")

	fmt.Println("Topic: " + topicSend + " will be used to writing messages.") // TODO
	fmt.Println("Topic: " + topicSold + " will be used to read messages.")

	connection(kafkaUrl, topicSold)
	go kaffka.StartKafkaReader(kafkaUrl, topicSold, topicSend)

	startApi()
}

func startApi() {
	r := chi.NewRouter()
	r.Get("/products/", handlers.GetAll)                                           // "Clients should be able to see a list of available products in the warehouse."
	r.Get("/products/{category}/{brand}/{maxPrice}", handlers.GetProductsFiltered) // "Clients should be able to get products based on product category, brand and maximum price"
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}

func connection(kafkaUrl string, topicName string) {
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
}

func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)

	if !ex {
		return defVal
	}

	return val
}
