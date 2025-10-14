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
	tokenSubString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", err
	}

	tokenString := "Bearer " + tokenSubString
	return tokenString, nil
}

func VerifyToken(tokenString string) (uuid.UUID, error) {

	tokenSubString := tokenString[7:]
	token, err := jwt.Parse(tokenSubString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenMalformed
		}
		return []byte(secret_key), nil
	})
	if err != nil {
		return uuid.Nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if customerID, ok := claims["customer_id"].(string); ok {
			parsedCustomerID, err := uuid.Parse(customerID)
			if err != nil {
				return uuid.Nil, err
			}
			return parsedCustomerID, nil
		}
	}
	return uuid.Nil, jwt.ErrTokenInvalidClaims
}
