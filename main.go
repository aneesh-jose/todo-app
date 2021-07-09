package main

import (
	"github.com/aneesh-jose/simple-server/controllers"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/", controllers.ReadTodos)
	app.Post("/", controllers.CreateTodo)
	app.Delete("/", controllers.DeleteTodo)
	app.Patch("/", controllers.UpdateTodo)
	app.Listen(8090)
}
