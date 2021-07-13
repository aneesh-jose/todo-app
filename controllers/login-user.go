package controllers

import (
	"database/sql"
	"fmt"

	"github.com/aneesh-jose/simple-server/models"
	"github.com/aneesh-jose/simple-server/packages/dbops"
	"github.com/gofiber/fiber"
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
	psqlInfo := dbops.GetDbCreds()
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

	defer db.Close()

	tokenString, err := JWTGenerator(body)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
	}
	ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"username": Username,
		"token":    tokenString,
	})

}
