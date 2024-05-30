package controllers

import (
	"encoding/json"
	"fer/src/application/usecases"
	"fer/src/domain/entities"
	"io"
	"net/http"
)

type CreateProductController struct {
	useCase usecases.CreateProductUseCase
}

func NewCreateProductController(useCase usecases.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{
		useCase: useCase,
	}
}

type ProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func (controller CreateProductController) Run(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var product ProductInput
	err = json.Unmarshal(body, &product)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}

	productEntity := entities.NewProduct(product.Name, product.Price, product.Stock)

	result, err := controller.useCase.Run(productEntity)
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"data":result})
}
