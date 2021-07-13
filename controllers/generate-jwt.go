package controllers

import (
	"time"

	"github.com/aneesh-jose/simple-server/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func JWTGenerator(body models.User) (string, error) {

	expirationTime := time.Now().Add(30 * 24 * time.Hour)

	claims := models.Claims{
		Username: body.Username,
		Password: body.Password,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	// Create the JWT string
	jwtKey := viper.Get("JWTKEY").(string)
	jsonKey := []byte(jwtKey)
	tokenString, err := token.SignedString(jsonKey)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", err
	}
	return tokenString, nil
}
