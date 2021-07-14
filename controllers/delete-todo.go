package controllers

import (
	"fmt"

	authentication "github.com/aneesh-jose/simple-server/utils/auth"
	"github.com/aneesh-jose/simple-server/utils/dbops"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func DeleteTodo(ctx *fiber.Ctx) {

	token := ctx.Cookies("token")
	//authenticate and generate username
	username, err := authentication.JWTAuthenticate(&token)
	if username == "" || err != nil {
		// jwt parsing/ authentication is error
		ctx.Status(fiber.StatusUnauthorized)
		return
	}

	type Todo struct {
		Id int `json:"id"`
	}
	var body Todo
	// parse json input from user to Todo struct
	err = ctx.BodyParser(&body)

	if err != nil {
		// improper input of json paramenters by the user
		// like type that does not match or absence of data fields
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request",
		})
		return
	}

	// obtain the id of the todo from the request
	// obtain username from the token(mentioned previously)
	// execute the delete command
	_, err = dbops.DeleteTodoFromDb(body.Id, username)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "database deletion unsuccessful",
		})
		fmt.Println(err)
		return
	}
	// send the 202(accepted) statuscode to the user
	// so that they know the deletion opetation has successfully completed
	ctx.Status(fiber.StatusAccepted)
}
