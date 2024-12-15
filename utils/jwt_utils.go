package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func CreateJWTToken(id string, role string) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = id
	claims["exp"] = exp
	claims["role"] = role
	t, err := token.SignedString([]byte(viper.GetString("app.jwt-secret")))

	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}