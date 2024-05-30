package usecases

import "fer/src/domain/entities"


type GetAllProductsUseCase struct {
	repository entities.IProduct
}

func NewGetAllProductsUseCase(repository entities.IProduct) *GetAllProductsUseCase {
	return &GetAllProductsUseCase{
		repository: repository,
	}
}

func (useCase *GetAllProductsUseCase) Run() ([]entities.Product, error) {
	return useCase.repository.GetAll()
}