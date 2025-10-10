package models

import (
	"example.com/goods-manage/db"
	"github.com/google/uuid"
)

type Category struct {
	CategoryID   uuid.UUID `json:"category_id"`
	CategoryName string    `json:"category_name"`
}

func GetAllCategories() ([]Category, error) {
	rows, err := db.DB.Query("SELECT category_id, category_name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.CategoryID, &category.CategoryName); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
