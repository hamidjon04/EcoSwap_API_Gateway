package token

import (
	"api_gateway/config"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func ValidToken(tokenString string) (bool, error) {
	_, err := ExtraClaims(tokenString)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtraClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Load().SIGNING_KEY), nil
	})
	
	if err != nil {
		return nil, fmt.Errorf("faieled to parse token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
