package usecases

import "fer/src/domain/entities"

type CreateProductUseCase struct {
	repository entities.IProduct
}

func NewProductUseCase(repository entities.IProduct) CreateProductUseCase {
	return CreateProductUseCase{
		repository: repository,
	}
}

func (useCase *CreateProductUseCase) Run(data entities.Product) (entities.Product, error) {
	return useCase.repository.Create(data)
}