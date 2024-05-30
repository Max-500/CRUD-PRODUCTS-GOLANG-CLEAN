package entities

import "github.com/google/uuid"

type IProduct interface {
	Create(Product) (Product, error)
	GetAll() ([]Product, error)
	Delete(uuid.UUID) (bool, error)
}