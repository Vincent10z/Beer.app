// users/repository/product_repository.go
package repository

import (
	"Beer.app/models"
	"errors"
)

type ProductRepository interface {
	GetProductById(id int) (*models.Product, error)
	CreateProduct(product *models.Product) error
}

type productRepository struct {
	products map[int]*models.Product
}

func NewProductRepository() ProductRepository {
	return &productRepository{products: make(map[int]*models.Product)}
}

func (r *productRepository) GetProductById(id int) (*models.Product, error) {
	product, exists := r.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	r.products[product.Id] = product
	return nil
}
