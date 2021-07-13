package controllers

import (
	"database/sql"
	"fmt"

	"github.com/aneesh-jose/simple-server/models"
	"github.com/aneesh-jose/simple-server/packages/dbops"
	"github.com/gofiber/fiber"
)

func CreateUser(ctx *fiber.Ctx) {

	var body models.UserDetails

	err := ctx.BodyParser(&body)
	if err != nil {

		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request",
		})
		return
	}

	psqlInfo := dbops.GetDbCreds()
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
	defer db.Close()
	ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}
