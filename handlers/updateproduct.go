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
	var neededProducts []models.ProductsRespSuppliers

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
			payload, err := json.Marshal(prodSupReq)

			// Create a new POST request with the specified URL and payload
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
			if err != nil {
				fmt.Println("Error creating POST request:", err)
				return
			}

			// Set the Content-Type header
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("accept", "application/json")

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

			var productsResSup models.ProdSupliers
			errUnm := json.Unmarshal(body.Bytes(), &productsResSup)
			if errUnm != nil {
				fmt.Println("Error:", errUnm)
				return
			}

			for _, productAdd := range productsResSup.Available.Products {
				// Adds the product to the DataBase => Update
				models.UpdateProductName(productAdd)
				// Add to the array of products added to the DB to send through kafka
				neededProducts = append(neededProducts, productAdd)
				fmt.Printf("Name: %s, Brand: %s, Category: %s, Quantity: %d, Price: %f\n", product.Name, product.Brand, product.Category, product.Quantity, product.Price)
			}
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

func writeMessageKafka(kafkaUrl, topicName string, neededProducts []models.ProductsRespSuppliers) {
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
