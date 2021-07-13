package controllers

import (
	"time"

	"github.com/aneesh-jose/simple-server/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func JWTGenerator(body models.User) (string, error) {
	// function used to generate a java web token according
	// to the username and password from the user

	// expiration time is set to 30 days from the generation of the token
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

	// the secure private key is obtained from the
	// .env file using viper and is converted to string
	jwtKey := viper.Get("JWTKEY").(string)
	jsonKey := []byte(jwtKey)                       //converting the obtained key to byte format
	tokenString, err := token.SignedString(jsonKey) //generate token

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", err
	}
	// return the newly generated token string with
	// no error(nil)
	return tokenString, nil
}
