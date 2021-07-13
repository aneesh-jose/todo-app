package controllers

import (
	"database/sql"
	"fmt"

	"github.com/aneesh-jose/simple-server/packages/dbops"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func DeleteTodo(ctx *fiber.Ctx) {

	token := ctx.Cookies("token")
	username, err := JWTAuthenticate(&token)
	if username == "" || err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return
	}

	type Todo struct {
		Id int `json:"id"`
	}
	var body Todo

	err = ctx.BodyParser(&body)

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
	_, err = db.Exec("delete from todos where id=$1 and username=$2", body.Id, username)
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
