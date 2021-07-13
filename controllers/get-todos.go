package controllers

import (
	"database/sql"
	"fmt"

	"github.com/aneesh-jose/simple-server/models"
	"github.com/aneesh-jose/simple-server/packages/dbops"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func ReadTodos(ctx *fiber.Ctx) {

	token := ctx.Cookies("token")
	username, err := JWTAuthenticate(&token)
	if username == "" || err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return
	}
	var testsamples []models.TodoJson

	psqlInfo := dbops.GetDbCreds()

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "connection unsuccessful",
		})
		fmt.Println(err.Error())
		return
	}
	todos, err := db.Query("select * from todos where username=$1", username)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "database query unsuccessful",
		})
		return
	}
	for todos.Next() {
		var sample models.TodoJson
		err = todos.Scan(&sample.Id, &sample.Name, &sample.Description, &sample.Status, &sample.User)
		if err != nil {
			ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "parsing unsuccessful",
			})
			return
		}
		testsamples = append(testsamples, sample)
	}

	todolist := make(map[int]fiber.Map)
	for _, elem := range testsamples {
		todolist[elem.Id] = fiber.Map{
			"name":        elem.Name,
			"description": elem.Description,
			"status":      elem.Status,
		}
	}
	ctx.Status(fiber.StatusFound).JSON(todolist)
	defer db.Close()
}
