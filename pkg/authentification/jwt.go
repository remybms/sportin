package authentification

import (
	"sportin/database/dbmodel"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func GenerateJWTToken(secret string, userEntry *dbmodel.UserEntry) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userEntry.ID,
		"exp": jwt.TimeFunc().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(secret))
}

func ValidateJWTToken(secret, tokenString string) (float64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"].(float64), nil
	}

	return 0, err
}
