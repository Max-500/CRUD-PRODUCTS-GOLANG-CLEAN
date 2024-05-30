package routes

import (
	"fer/src/infraestructure"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	dependencies := infraestructure.NewDependencies()

	r.HandleFunc("/api/v1/products", dependencies.CreateProductController.Run).Methods("POST")
	r.HandleFunc("/api/v1/products", dependencies.GetAllProductController.Run).Methods("GET")
	r.HandleFunc("/api/v1/products/{uuid}", dependencies.DeleteProductController.Run).Methods("DELETE")

	return r
}
