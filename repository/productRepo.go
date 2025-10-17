package repository

import "example.com/goods-manage/models"

type ProductRepo interface {
	GetAllProducts() ([]models.Product, error)
}

type productRepo struct{}

func NewProductRepo() ProductRepo {
	return &productRepo{}
}

func (r *productRepo) GetAllProducts() ([]models.Product, error) {
	var prodModel models.Product
	return prodModel.GetAllProducts()
}