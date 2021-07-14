package controllers

import (
	"fmt"

	"github.com/aneesh-jose/simple-server/models"
	"github.com/aneesh-jose/simple-server/utils/dbops"
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

	result, err := dbops.AddUserToDb(body.Username, body.Password, body.Name)
	if err != nil {
		// this error may occur as the user is already signed up
		// and the primary key username throws error
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
