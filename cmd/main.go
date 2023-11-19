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

	fmt.Println("Entrei aqui")

	err := configs.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Get("/", handlers.GetAll)
	r.Delete("/{id}", handlers.Delete)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
