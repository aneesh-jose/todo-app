package controllers

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func CreateTodo(ctx *fiber.Ctx) {

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
	result, err := db.Exec("insert into todos values(nextval('countsequence'),$1,$2,$3)", body.Name, body.Description, false)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "database insertion unsuccessful",
		})
		fmt.Println(err)
		return
	}
	lastId, _ := result.LastInsertId()
	ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"accepted": lastId,
	})
}
