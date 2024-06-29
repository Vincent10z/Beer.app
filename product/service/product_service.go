// users/service/reviews_service.go
package service

import (
	"Beer.app/models"
	"Beer.app/product/repository"
)

type ProductService interface {
	GetProduct(id int) (*models.Product, error)
	CreateProduct(user *models.Product) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetProduct(id int) (*models.Product, error) {
	return s.repo.GetProductById(id)
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.repo.CreateProduct(product)
}
