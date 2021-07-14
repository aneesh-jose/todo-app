package controllers

import (
	"database/sql"
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

	psqlInfo := dbops.GetDbCreds() //DB credentials in string format
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		// credentials are not authentic
		// may be because the parameters are false or
		// The database server may be down and need to restart
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "database connection unsuccessful",
		})
		fmt.Println(err.Error())
		return
	}
	// insert the new user into the 'USERS' table
	result, err := db.Exec("insert into users values($1,$2,$3)", body.Username, body.Password, body.Name)
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
	defer db.Close()
	ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}
