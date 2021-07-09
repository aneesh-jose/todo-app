package main

import (
	"github.com/aneesh-jose/simple-server/controllers"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	// todo addition and deletion
	app.Get("/", controllers.ReadTodos)
	app.Post("/", controllers.CreateTodo)
	app.Delete("/", controllers.DeleteTodo)
	app.Patch("/", controllers.UpdateTodo)

	// user creation
	app.Post("/signup", controllers.CreateUser)
	app.Post("/login", controllers.Login)

	app.Listen(8090)
}
