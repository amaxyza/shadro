package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int, username string) (string, error) {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	return "", err
	// }

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"id":       id,
			"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	)

	token_str, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token_str, nil
}

func VerifyToken(jwt_str string) (string, int, error) {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	return "", -1, err
	// }

	token, err := jwt.Parse(
		jwt_str,
		func(token *jwt.Token) (any, error) { return []byte(os.Getenv("JWT_SECRET")), nil },
	)

	if err != nil {
		return "", -1, err
	}

	if !token.Valid {
		return "", -1, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username, ok := claims["username"].(string)
		if !ok {
			return "", -1, fmt.Errorf("username not found in token")
		}
		id, ok := claims["id"].(float64)
		if !ok {
			return "", -1, fmt.Errorf("id not found in token")
		}

		return username, int(id), nil
	}

	return "", -1, fmt.Errorf("invalid claims")
}
