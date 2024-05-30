package controllers

import (
	"encoding/json"
	"fer/src/application/usecases"
	"net/http"
)

type GetAllProductsController struct {
	useCase *usecases.GetAllProductsUseCase
}

func NewGetAllProductsController(useCase *usecases.GetAllProductsUseCase) *GetAllProductsController {
	return &GetAllProductsController{
		useCase: useCase,
	}
}

func (controller *GetAllProductsController) Run(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := controller.useCase.Run()
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"data":result})
}
