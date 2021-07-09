package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func JWTAuthenticate(token *string, username *string) (bool, error) {

	type Claims struct {
		Username string `json:"username"`
		Password string `json:"password"`
		jwt.StandardClaims
	}

	claims := &Claims{}

	jwtKey := viper.Get("JWTKEY").(string)
	jsonKey := []byte(jwtKey)

	tkn, err := jwt.ParseWithClaims(*token, claims, func(token *jwt.Token) (interface{}, error) {
		return jsonKey, nil
	})
	if err != nil {
		return false, err
	}
	if !tkn.Valid {

		return false, err
	}
	return true, nil
}
