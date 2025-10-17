package service

import (
	"example.com/goods-manage/models"
	"example.com/goods-manage/repository"
)

type ProductService struct {
	productRepo repository.IProductRepo
}

func NewProductService(productRepo repository.IProductRepo) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.productRepo.GetAllProducts()
}