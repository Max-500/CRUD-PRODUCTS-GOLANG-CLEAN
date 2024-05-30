package usecases

import (
	"fer/src/domain/entities"

	"github.com/google/uuid"
)

type DeleteProductUseCase struct {
	repository entities.IProduct
}

func NewDeleteProductUseCase(repository entities.IProduct) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		repository: repository,
	}
}

func (useCase *DeleteProductUseCase) Run(uuid uuid.UUID) (bool, error) {
	return useCase.repository.Delete(uuid)
}