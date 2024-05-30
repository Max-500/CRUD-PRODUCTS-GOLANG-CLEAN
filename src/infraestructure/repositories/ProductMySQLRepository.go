package repositories

import (
	mysql "fer/Database/MySQL"
	"fer/src/domain/entities"

	"github.com/google/uuid"
)

type ProductMySQLRepository struct {
	db *mysql.DatabaseMySQL
}

func NewProductMySQLRepository(db *mysql.DatabaseMySQL) *ProductMySQLRepository {
	return &ProductMySQLRepository{
		db: db,
	}
}

func (repo *ProductMySQLRepository) Create(product entities.Product) (entities.Product, error) {
	query := `INSERT INTO products (uuid, name, price, stock) VALUES (?, ?, ?, ?)`
	_, err := repo.db.Conn.Exec(query, product.Uuid, product.Name, product.Price, product.Stock)
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (repo *ProductMySQLRepository) GetAll() ([]entities.Product, error) {
	query := "SELECT * FROM products;"
	rows, err := repo.db.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		err = rows.Scan(&product.Uuid, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *ProductMySQLRepository) Delete(uuid uuid.UUID) (bool, error) {
	query := "DELETE FROM products WHERE uuid = ? LIMIT 1; "
	result, err := repo.db.Conn.Exec(query, uuid)

	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}