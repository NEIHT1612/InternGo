package repository

import "example.com/goods-manage/models"

type IProductRepo interface {
	GetAllProducts() ([]models.Product, error)
}

type productRepo struct{}

func NewProductRepo() IProductRepo {
	return &productRepo{}
}

func (r *productRepo) GetAllProducts() ([]models.Product, error) {
	var prodModel models.Product
	return prodModel.GetAllProducts()
}
