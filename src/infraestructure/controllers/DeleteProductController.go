package controllers

import (
	"encoding/json"
	"fer/src/application/usecases"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type DeleteProductController struct {
	useCase *usecases.DeleteProductUseCase
}

func NewDeleteProductController(useCase *usecases.DeleteProductUseCase) *DeleteProductController {
	return &DeleteProductController{
		useCase: useCase,
	}
}

func (controller *DeleteProductController) Run(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	uuidParam, ok := params["uuid"]
	if !ok {
		http.Error(w, "UUID not provided", http.StatusBadRequest)
		return
	}

	uuid, err := uuid.Parse(uuidParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	result, err := controller.useCase.Run(uuid)
	if err != nil || !result {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"data":"Products was destroy successfully"})
}
