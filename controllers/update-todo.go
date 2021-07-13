package controllers

import (
	"database/sql"
	"fmt"

	"github.com/aneesh-jose/simple-server/packages/dbops"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func UpdateTodo(ctx *fiber.Ctx) {

	token := ctx.Cookies("token")
	username, err := JWTAuthenticate(&token)
	if username == "" || err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return
	}

	type Todo struct {
		Id     int  `json:"id"`
		Status bool `json:"status"`
	}
	var body Todo

	err = ctx.BodyParser(&body)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Cannot parse request",
			"details": err.Error(),
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
	result, err := db.Exec("update todos set status=$1 where id=$2 and username=$3", body.Status, body.Id, username)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "database updation unsuccessful",
		})
		fmt.Println(err)
		return
	}
	lastId, _ := result.LastInsertId()
	ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"accepted": lastId,
	})
}
