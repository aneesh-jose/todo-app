package controllers

import (
	"github.com/aneesh-jose/simple-server/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func JWTAuthenticate(token *string) (string, error) {

	claims := models.Claims{}

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	jwtKey := viper.Get("JWTKEY").(string)
	jsonKey := []byte(jwtKey)

	tkn, err := jwt.ParseWithClaims(*token, &claims, func(token *jwt.Token) (interface{}, error) {
		return jsonKey, nil
	})
	if err != nil {
		return "", err
	}
	if !tkn.Valid {

		return "", err
	}
	return claims.Username, nil
}
