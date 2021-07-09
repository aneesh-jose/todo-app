package controllers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func CreateTodo(ctx *fiber.Ctx) {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	host := viper.Get("HOST")
	user := viper.Get("USER")
	password := viper.Get("PASSWORD")
	dbname := viper.Get("DBNAME")
	portStr, _ := viper.Get("PORT").(string)
	port, _ := strconv.Atoi(portStr)

	type Todo struct {
		Name        string `json:"name"`
		Description string `json:"description"`
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
	result, err := db.Exec("insert into todos values(nextval('countsequence'),$1,$2,$3)", body.Name, body.Description, false)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "database insertion unsuccessful",
		})
		fmt.Println(err)
		return
	}
	lastId, _ := result.LastInsertId()

	defer db.Close()

	ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"accepted": lastId,
	})
}
