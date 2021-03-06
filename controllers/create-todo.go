package controllers

import (
	authentication "github.com/aneesh-jose/simple-server/utils/auth"
	"github.com/aneesh-jose/simple-server/utils/dbops"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func CreateTodo(ctx *fiber.Ctx) {

	token := ctx.Cookies("token")
	username, err := authentication.JWTAuthenticate(&token)
	if username == "" || err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return
	}

	// The format of user input
	type Todo struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	var body Todo

	err = ctx.BodyParser(&body)

	if err != nil {
		// most probably the user input does not satisfy the
		// preferred format
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request",
		})
		return
	}
	// insert the newly obtained todo to the 'TODOS' table
	// the name and the description of the todo is obtained from user end
	// the status of completion of the todo is marked as false(not completed) and
	// username is obtained from the jwt authenticator
	// result, err := db.Exec("insert into todos values(nextval('countsequence'),$1,$2,$3,$4)", body.Name, body.Description, false, username)
	result, err := dbops.AddTodoToDB(body.Name, body.Description, username)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}
	lastId, _ := result.LastInsertId()

	ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"accepted": lastId,
	})
}
