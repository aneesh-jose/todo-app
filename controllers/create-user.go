package controllers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/spf13/viper"
)

func CreateUser(ctx *fiber.Ctx) {

	type User struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	var body User

	err := ctx.BodyParser(&body)
	if err != nil {

		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request",
		})
		return
	}

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

	result, err := db.Exec("insert into users values($1,$2,$3)", body.Username, body.Password, body.Name)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
		fmt.Println(err)
		return
	}
	id, err := result.LastInsertId()
	ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}
