package controllers

import (
	"github.com/aneesh-jose/simple-server/models"
	authentication "github.com/aneesh-jose/simple-server/utils/auth"
	"github.com/aneesh-jose/simple-server/utils/dbops"
	"github.com/gofiber/fiber"
)

func Login(ctx *fiber.Ctx) {

	var body models.User
	err := ctx.BodyParser(&body) // parse the input data to model.User structure
	if err != nil {
		// the user input is wrong
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request",
		})
		return
	}

	var Username string

	// query the table 'USERS' on username and password from the user
	query, err := dbops.CheckUserAvailability(body)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "database query unsuccessful",
		})
		return
	}

	for query.Next() {
		err = query.Scan(&Username) //add to username
		if err != nil {
			ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "parsing unsuccessful",
			})
			return
		}
	}
	// user not found
	if Username == "" {
		ctx.Status(fiber.StatusUnauthorized)
	}

	// generate token according to the username and password
	tokenString, err := authentication.JWTGenerator(body)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
	}
	// send the username and token to user
	ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"username": Username,
		"token":    tokenString,
	})

}
