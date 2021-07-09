package controllers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/aneesh-jose/simple-server/models"
	"github.com/gofiber/fiber"
	"github.com/spf13/viper"
)

func Login(ctx *fiber.Ctx) {

	var body models.User
	err := ctx.BodyParser(&body)
	if err != nil {

		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request",
		})
		return
	}

	var Username string

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	host := viper.Get("HOST")
	user := viper.Get("USER")
	password := viper.Get("PASSWORD")
	dbname := viper.Get("DBNAME")
	portStr, _ := viper.Get("PORT").(string)
	port, _ := strconv.Atoi(portStr)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "connection unsuccessful",
		})
		fmt.Println(err.Error())
		return
	}
	query, err := db.Query("select username from users where username=$1 and password=$2", body.Username, body.Password)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "database query unsuccessful",
		})
		return
	}

	for query.Next() {
		err = query.Scan(&Username)
		if err != nil {
			ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "parsing unsuccessful",
			})
			return
		}
	}
	if Username == "" {
		ctx.Status(fiber.StatusUnauthorized)
	}
	tokenString, err := JWTGenerator(body)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
	}
	ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"username": Username,
		"token":    tokenString,
	})

}
