package controllers

import (
	"fmt"

	authentication "github.com/aneesh-jose/simple-server/utils/auth"
	"github.com/aneesh-jose/simple-server/utils/dbops"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func UpdateTodo(ctx *fiber.Ctx) {

	token := ctx.Cookies("token")                           // read token from cookies
	username, err := authentication.JWTAuthenticate(&token) // authenticate the token
	if username == "" || err != nil {
		// error while parsing the jwt
		ctx.Status(fiber.StatusUnauthorized)
		return
	}

	type Todo struct {
		Id     int  `json:"id"`
		Status bool `json:"status"`
	}
	var body Todo

	// Parse the input from user.
	// Its content will be the id of the todo and the status it need to change to(completed or not)
	err = ctx.BodyParser(&body)

	if err != nil {
		// Illegal request: the request body cannot be parsed
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Cannot parse request",
			"details": err.Error(),
		})
		return
	}

	// update the todo data i.e, the status (completed or not) according to the user input
	result, err := dbops.UpdateTodoOperation(body.Id, body.Status)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "database updation unsuccessful",
		})
		fmt.Println(err)
		return
	}
	// successfully updated the database
	lastId, _ := result.LastInsertId()
	ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"accepted": lastId,
	})
}
