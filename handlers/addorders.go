package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tamiresviegas/warehouse/models"
)

func AddOrder(w http.ResponseWriter, r *http.Request) {

	var orderP models.OrderP
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&orderP)
	if err != nil {
		log.Println("Erro ao fazer decode do json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	id, err := models.AddOrderP(orderP)
	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("INSERIDO COM SUCESSO! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
