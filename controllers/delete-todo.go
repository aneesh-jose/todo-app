package controllers

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func DeleteTodo(ctx *fiber.Ctx) {
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

	user := "postgres"
	password := "root"
	dbname := "postgres"
	port := 5432
	host := "localhost"

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

	ctx.Status(fiber.StatusAccepted)
}
