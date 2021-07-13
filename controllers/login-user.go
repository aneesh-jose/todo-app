package controllers

import (
	"database/sql"

	"github.com/aneesh-jose/simple-server/models"
	"github.com/aneesh-jose/simple-server/packages/dbops"
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
	psqlInfo := dbops.GetDbCreds()            //obtain database credentials
	db, err := sql.Open("postgres", psqlInfo) //connect to database
	if err != nil {
		// database connection error
		// maybe database server is down or the
		// authentication credentials might have changed
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "connection unsuccessful",
		})
		return
	}
	// query the table 'USERS' on username and password from the user
	query, err := db.Query("select username from users where username=$1 and password=$2", body.Username, body.Password)
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

	defer db.Close()
	// generate token according to the username and password
	tokenString, err := JWTGenerator(body)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
	}
	// send the username and token to user
	ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"username": Username,
		"token":    tokenString,
	})

}
