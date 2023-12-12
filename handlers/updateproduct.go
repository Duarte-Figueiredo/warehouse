package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/segmentio/kafka-go"
	"github.com/tamiresviegas/warehouse/models"
)

// Receives an array of products and updates them in the database
func UpdateProducts(kafkaUrl string, message kafka.Message, topicSend string) {
	fmt.Println("receive a message: ", string(message.Value)) // [{"product_id":,"quantity":3},{"product_id":"product2","quantity":2},{"product_id":"product3","quantity":1}]

	var prodQuantites []models.ProductQntUpdt
	var neededProducts []models.Product

	if err := json.Unmarshal([]byte(string(message.Value)), &prodQuantites); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	for i := 0; i < len(prodQuantites); i++ {

		if prodQuantites[i].Quantity == 0 {
			// Asks more of the produtct for supliers

			// Gets the product with the specified id
			product, err := models.GetProduct(prodQuantites[i].Product_ID)
			if err != nil {
				fmt.Println("Error while trying to get the product ", err)
				return
			}
			fmt.Println("Received: ", product.Product_ID, product.Brand, product.Category, product.Name, product.Price, product.Quantity)

			var prodReq models.ProdReq
			prodReq.Category = product.Category
			prodReq.Price = product.Price
			prodReq.Quantity = product.Quantity

			var prodReqList []models.ProdReq
			var prodSupReq models.ProductSuppliersReq

			prodReqList = append(prodReqList, prodReq)
			prodSupReq.Products = prodReqList

			// Makes the product request
			url := "http://34.16.134.101:8081/orders"
			//payload, err := json.Marshal(prodSupReq)

			//fmt.Println("Payload: ", string(payload))

			payload := []byte(`{"products":[{"category":"Rice","quantity":3,"price":1.4}]}`)
			// Create a new POST request with the specified URL and payload
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
			if err != nil {
				fmt.Println("Error creating POST request:", err)
				return
			}

			// Set the Content-Type header
			req.Header.Set("Content-Type", "application/json")

			// Perform the request
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Println("Error executing POST request:", err)
				return
			}
			defer resp.Body.Close()

			// Check the HTTP status code
			fmt.Println("Status Code:", resp.Status)

			// Read and print the response body
			body := new(bytes.Buffer)
			_, err = body.ReadFrom(resp.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)
				return
			}

			fmt.Println("Response Body:", body.String())

			// Adds the product to the DB

			// Sends a message through kafka saying a product was updated
			/*var product models.Product
			product.Product_ID = 1
			product.Name = "testeJen22"
			product.Brand = "teste2"
			product.Category = "teste3"
			product.Quantity = 4
			product.Price = 3
			neededProducts = append(neededProducts, product)*/
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

func writeMessageKafka(kafkaUrl, topicName string, neededProducts []models.Product) {
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

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(jsonData),
	})

	if err != nil {
		log.Fatal("cannot write a message: ", err)
	}
}
