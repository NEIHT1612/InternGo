package service

import (
	"example.com/goods-manage/models"
	"example.com/goods-manage/repository"
)

type ProductService struct {
	productRepo repository.ProductRepo
}

func NewProductService(productRepo repository.ProductRepo) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.productRepo.GetAllProducts()
}