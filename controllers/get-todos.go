package controllers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/aneesh-jose/simple-server/models"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func ReadTodos(ctx *fiber.Ctx) {

	token := ctx.Cookies("token")
	username, err := JWTAuthenticate(&token)
	if username == "" || err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return
	}
	fmt.Println(username)
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	host := viper.Get("HOST")
	user := viper.Get("USER")
	password := viper.Get("PASSWORD")
	dbname := viper.Get("DBNAME")
	portStr, _ := viper.Get("PORT").(string)
	port, _ := strconv.Atoi(portStr)

	var testsamples []models.TodoJson

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "connection unsuccessful",
		})
		fmt.Println(err.Error())
		return
	}
	query, err := db.Query("select * from todos where username=$1", username)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "database query unsuccessful",
		})
		return
	}

	for query.Next() {
		var sample models.TodoJson
		err = query.Scan(&sample.Id, &sample.Name, &sample.Description, &sample.Status, &sample.User)
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
