package infraestructure

import (
	mysql "fer/Database/MySQL"
	"fer/src/application/usecases"
	"fer/src/infraestructure/controllers"
	"fer/src/infraestructure/repositories"
)

type Dependencies struct {
	CreateProductController *controllers.CreateProductController
	GetAllProductController *controllers.GetAllProductsController
	DeleteProductController *controllers.DeleteProductController
}

func NewDependencies() *Dependencies {
	db, err := mysql.NewDatabase("root", "", "localhost", "mydatabase")
	if err != nil {
		panic(err)
	}

	repo := repositories.NewProductMySQLRepository(db)
	createProductUseCase := usecases.NewProductUseCase(repo)

	getAllProductUseCase := usecases.NewGetAllProductsUseCase(repo)

	deleteProductUseCase := usecases.NewDeleteProductUseCase(repo)

	return &Dependencies{
		CreateProductController: controllers.NewCreateProductController(createProductUseCase),
		GetAllProductController: controllers.NewGetAllProductsController(getAllProductUseCase),
		DeleteProductController: controllers.NewDeleteProductController(deleteProductUseCase),
	}
}