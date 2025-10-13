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

func GetCategoryByID(categoryID uuid.UUID) (*Category, error) {
	row := db.DB.QueryRow("SELECT category_id, category_name FROM categories WHERE category_id = $1", categoryID)
	var category Category
	if err := row.Scan(&category.CategoryID, &category.CategoryName); err != nil {
		return nil, err
	}
	return &category, nil
}

func CreateCategory(category *Category) error {
	query := `INSERT INTO categories (category_name) VALUES ($1) RETURNING category_id`
	err := db.DB.QueryRow(query, category.CategoryName).Scan(&category.CategoryID)
	if err != nil {
		return err
	}	
	return nil
}

func UpdateCategoryByID(category *Category) error {
	query := `UPDATE categories SET category_name = $1 WHERE category_id = $2 RETURNING category_id, category_name`
	err := db.DB.QueryRow(query, category.CategoryName, category.CategoryID).Scan(&category.CategoryID, &category.CategoryName)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategoryByID(categoryID uuid.UUID) error {
	query := `DELETE FROM categories WHERE category_id = $1 RETURNING category_id`
	err := db.DB.QueryRow(query, categoryID).Scan(&categoryID)
	if err != nil {
		return err
	}
	return nil
}
