package main

import (
	"fmt"
	"net/http"

	"github.com/tamiresviegas/warehouse/handlers"

	"github.com/go-chi/chi"
	"github.com/tamiresviegas/warehouse/configs"
)

func main() {
	err := configs.Load()
	if err != nil {
		return
	}

	r := chi.NewRouter()

	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.GetAll)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}


