package controllers

import (
	"github.com/aneesh-jose/simple-server/models"
	authentication "github.com/aneesh-jose/simple-server/utils/auth"
	"github.com/aneesh-jose/simple-server/utils/dbops"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func ReadTodos(ctx *fiber.Ctx) {
	// function used to send user the todos generatd by them

	token := ctx.Cookies("token")                           //obtain the token from cookies
	username, err := authentication.JWTAuthenticate(&token) //authenticate the token
	if username == "" || err != nil {
		// if the parsing of the token is error
		// the user is not valid and therefore send
		// an unauthorized status
		ctx.Status(fiber.StatusUnauthorized)
		return
	}

	var todoList []models.TodoJson

	todos, err := dbops.ReadTodosFromDb(username)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "database query unsuccessful",
			"stat":  err.Error(),
		})
		return
	}
	for todos.Next() {
		// traverse through all the results
		var sample models.TodoJson
		// and parse them to the standard TODO format
		err = todos.Scan(&sample.Id, &sample.Name, &sample.Description, &sample.Status, &sample.User)
		if err != nil {
			// the data from the database is corrupt
			// or may be an internal server error occured while parsing
			ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "parsing unsuccessful",
			})
			return
		}
		// add the the parsed ones to the todo list
		todoList = append(todoList, sample)
	}
	// todo list to be the output to the user
	// in the format id:{tododetails..} in fiber.Map format
	// fiber.Map is used insted of normal map is because its value is an interface
	// and therefore can contain any representable datatypes
	outputTodoMap := make(map[int]fiber.Map)

	for _, elem := range todoList {
		outputTodoMap[elem.Id] = fiber.Map{
			"name":        elem.Name,
			"description": elem.Description,
			"status":      elem.Status,
		}
	}
	// send the data with status as found
	ctx.Status(fiber.StatusFound).JSON(outputTodoMap)
}
