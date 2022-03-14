package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

func (factory *MiddlewareFactory) NewJwtAuthMiddleware(bearerHeader string) (interface{}, error) {
	bearerToken := strings.Split(bearerHeader, " ")[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error decoding token")
		}
		return []byte("12345"), nil
	})
	if err != nil {
		factory.logger.Error("", err, nil)
		return nil, err
	}

	if token.Valid {
		return claims["user"].(string), nil
	}
	return nil, errors.New("invalid token")
}
