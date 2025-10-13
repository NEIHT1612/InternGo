package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const secret_key = "justmeknow"

func GenerateToken(Username string, CustomerID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":    Username,
		"customer_id": CustomerID,
		"exp":         time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
