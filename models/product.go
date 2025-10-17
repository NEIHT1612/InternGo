package models

import (
	"example.com/goods-manage/db"
	"github.com/google/uuid"
)

type Product struct {
	ProductID     uuid.UUID 	`json:"product_id"`
	ProductName   string    	`json:"product_name"`
	CategoryName  string    	`json:"category_name"`
	SupplierName  string    	`json:"supplier_name"`
	Unit          string    	`json:"unit"`
	UnitPrice     float64    	`json:"unit_price"`
}

func (p *Product) GetAllProducts() ([]Product, error) {
	rows, err := db.DB.Query(`
		SELECT p.product_id, p.product_name, c.category_name, s.supplier_name, p.unit, p.unit_price
		FROM products p
		JOIN categories c ON p.category_id = c.category_id
		JOIN suppliers s ON p.supplier_id = s.supplier_id
	`)	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var prod Product
		err := rows.Scan(&prod.ProductID, &prod.ProductName, &prod.CategoryName, &prod.SupplierName, &prod.Unit, &prod.UnitPrice)
		if err != nil {
			return nil, err
		}
		products = append(products, prod)
	}
	return products, nil
}