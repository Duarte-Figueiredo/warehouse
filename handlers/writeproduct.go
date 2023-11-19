package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tamiresviegas/warehouse/models"
)

func Create(w http.ResponseWriter, r *http.Request) {

	var produto models.Produto
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		log.Println("Erro ao fazer decode do json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	id, err := models.Insert(produto)
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
