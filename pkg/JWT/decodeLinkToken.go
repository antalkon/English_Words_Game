package jwt

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type TokenRefreshClaims struct {
	ID      string `json:"userDisplay"`
	Email   string `json:"userEmail"`
	Phone   string `json:"userPhone"`
	Name    string `json:"userName"`
	Surname string `json:"userSurname"`
	Exp     int64  `json:"exp"`
	jwt.StandardClaims
}

func DecodeRefreshToken(refreshToken string) (*TokenRefreshClaims, error) {
	config, err := ReadConfigJWT()
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения конфигурации: %v", err)
	}

	token, err := jwt.ParseWithClaims(refreshToken, &TokenRefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("не удалось декодировать токен: %v", err)
	}

	if claims, ok := token.Claims.(*TokenRefreshClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("неверный токен")
	}
}
