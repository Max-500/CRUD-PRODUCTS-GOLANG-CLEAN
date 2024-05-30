package entities

import "github.com/google/uuid"

type Product struct {
	Uuid  uuid.UUID  `json:"uuid"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func NewProduct(name string, price float64, stock int) Product {
	return Product{
		Uuid:  uuid.New(),
		Name:  name,
		Price: price,
		Stock: stock,
	}
}