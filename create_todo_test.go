package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aneesh-jose/simple-server/controllers"
	"github.com/gofiber/fiber"
)

func TestCreateTodo(t *testing.T) {
	app := fiber.New()
	app.Post("/", func(ctx *fiber.Ctx) {
		token := ctx.Cookies("token")
		username, err := controllers.JWTAuthenticate(&token)
		if username == "" || err != nil {
			ctx.Status(fiber.StatusUnauthorized)
			return
		}
		type Todo struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		}
		var body Todo

		err = ctx.BodyParser(&body)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse request",
			})
			return
		}

		ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"accepted": 0,
		})
	})

	testBody := []struct {
		name   string
		body   map[string]string
		ctype  string
		cookie string
		resp   map[string]int
		code   int
	}{
		{
			name: "Accoring to requirements",
			body: map[string]string{"name": "the todo name3",
				"description": "The TODO description"},
			ctype:  "application/json",
			cookie: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFuZWVzaGpvc2UiLCJwYXNzd29yZCI6InNhbXBsZXBhc3MiLCJleHAiOjE2Mjg2NjU1Mjd9.LU90k2_qfJLINzbZbS5VkQre-5AIWlS6O2s6RAhOx4k",
			resp: map[string]int{
				"accepted": 0,
			},
			code: fiber.StatusAccepted,
		},
		{
			name: "Unauthorized user",
			body: map[string]string{"name": "the todo name3",
				"description": "The TODO description"},
			ctype:  "application/json",
			cookie: "eyJhbGciPiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFuZWVzaGpvc2UiLCJwYXNzd29yZCI6InNhbXBsZXBhc3MiLCJleHAiOjE2Mjg2NjU1Mjd9.LU90k2_qfJLINzbZbS5VkQre-5AIWlS6O2s6RAhOx4k",
			resp: map[string]int{
				"accepted": 0,
			},
			code: fiber.StatusUnauthorized,
		},
	}

	for _, post := range testBody {

		body, _ := json.Marshal(post.body)

		req, err := http.NewRequest("POST", "http://localhost:8090/", bytes.NewReader(body))

		if err != nil {
			t.Errorf("Name: %v, Error on request:%v", post.name, err)
		}

		req.AddCookie(&http.Cookie{Name: "token", Value: post.cookie})
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		if err != nil {
			t.Errorf("Name: %v, Error on request:%v", post.name, err)
		}

		if resp.StatusCode != post.code {
			t.Errorf("Name: %v, Expected statuscode %v, received:%v", post.name, post.code, resp.StatusCode)
		}
	}
}
