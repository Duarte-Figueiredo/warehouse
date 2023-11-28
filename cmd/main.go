package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tamiresviegas/warehouse/configs"

	"github.com/go-chi/chi"
	"github.com/tamiresviegas/warehouse/handlers"
)

func main() {

	fmt.Println("Started warehouse app")

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
