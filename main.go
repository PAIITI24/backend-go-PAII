package main

import (
	"PAII_Back-End/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable: false,
		AppName:   "Aplikasi Apotek Toba",
	})

	app.Post("/user/create", controllers.CreateUser)
	app.Listen(":8080")
}
