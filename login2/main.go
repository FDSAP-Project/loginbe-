package main

import (
	"sample/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.InitialMigration()
	app := fiber.New()

	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("/register", Register)
	app.Post("/login", Login)
	app.Listen(":3000")
}

type Log struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var user database.User
	c.BodyParser(&user)
	database.Db.Create(&user)
	return c.JSON(&fiber.Map{
		"message": "user registerd",
		"data":    user,
	})
}

func Login(c *fiber.Ctx) error {
	var log Log
	var user database.User
	c.BodyParser(&log)
	database.Db.Raw("select * from users where username=? and password=?", log.Username, log.Password).Find(&user)
	if user.Username == "" && user.Password == "" {
		return c.JSON(&fiber.Map{
			"message": "loggin failed",
		})
	}

	return c.JSON(&fiber.Map{
		"message":  "loggin success",
		"Welcome ": user.Name,
	})
}
