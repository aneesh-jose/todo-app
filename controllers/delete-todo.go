package controllers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func DeleteTodo(ctx *fiber.Ctx) {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	host := viper.Get("HOST")
	user := viper.Get("USER")
	password := viper.Get("PASSWORD")
	dbname := viper.Get("DBNAME")
	portStr, _ := viper.Get("PORT").(string)
	port, _ := strconv.Atoi(portStr)

	type Todo struct {
		Id int `json:"id"`
	}
	var body Todo

	err := ctx.BodyParser(&body)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request",
		})
		return
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "connection unsuccessful",
		})
		fmt.Println(err.Error())
		return
	}
	_, err = db.Exec("delete from todos where id=$1", body.Id)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "database deletion unsuccessful",
		})
		fmt.Println(err)
		return
	}
	defer db.Close()

	ctx.Status(fiber.StatusAccepted)
}
