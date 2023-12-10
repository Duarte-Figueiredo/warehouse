package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
	"github.com/tamiresviegas/warehouse/models"
)

// Receives an array of products and updates them in the database
func UpdateProducts(message kafka.Message, kafkaUrl string, topicSend string) {
	fmt.Println("receive a message: ", string(message.Value)) // [{"product_id":,"quantity":3},{"product_id":"product2","quantity":2},{"product_id":"product3","quantity":1}]

	var prodQuantites []models.ProductQntUpdt
	var neededProducts []models.Product

	if err := json.Unmarshal([]byte(string(message.Value)), &prodQuantites); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	for i := 0; i < len(prodQuantites); i++ {

		if prodQuantites[i].Quantity == 0 {
			//  asks more of the produtct for supliers

			// Sends a message through kafka saying a product was updated
			var product models.Product
			product.Product_ID = 1
			product.Name = "teste"
			product.Brand = "teste2"
			product.Category = "teste3"
			product.Quantity = 4
			product.Price = 3
			neededProducts = append(neededProducts, product)
		} else {
			// Updates the products of the DB
			models.UpdateProduct(prodQuantites[i].Product_ID, prodQuantites[i].Quantity)
		}
	}

	if len(neededProducts) != 0 {
		writeMessageKafka(kafkaUrl, topicSend, neededProducts)
	}

	neededProducts = nil
	prodQuantites = nil
}

func writeMessageKafka(kafkaUrl string, topicName string, neededProducts []models.Product) {
	writer := &kafka.Writer{
		Addr:  kafka.TCP(kafkaUrl),
		Topic: topicName,
	}

	// Convert struct to JSON
	jsonData, errJson := json.Marshal(neededProducts)
	if errJson != nil {
		fmt.Println("Error: ", errJson)
		return
	}

	fmt.Println("Writing test message to topic " + topicName)
	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(jsonData),
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
}
