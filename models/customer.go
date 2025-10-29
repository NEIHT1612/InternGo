package models

import (
	"errors"

	"example.com/goods-manage/db"
	"example.com/goods-manage/common"
	"github.com/google/uuid"
)

type Customer struct {
	CustomerID uuid.UUID `json:"customer_id"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
}

func CreateUser(customer *Customer) error {
	query := `
        INSERT INTO goods_management.customers (username, password_hash)
        VALUES ($1, $2)
        RETURNING customer_id, username
    `
	password_hash, err := common.HashPassword(customer.Password)
	if err != nil {
		return err
	}

	err = db.DB.QueryRow(query, customer.Username, password_hash).Scan(&customer.CustomerID, &customer.Username)
	if err != nil {
		return err
	}
	return nil
}

func LoginUser(customer *Customer) error {
	query := `
		SELECT customer_id, username, password_hash
		FROM customers
		WHERE username = $1
	`
	var password_hash string
	err := db.DB.QueryRow(query, customer.Username).Scan(&customer.CustomerID, &customer.Username, &password_hash)
	if err != nil {
		return err
	}
	if !common.CheckPasswordHash(customer.Password, password_hash) {
		return errors.New("invalid credentials")
	}
	return nil
}
