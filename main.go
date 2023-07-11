package main

import (
	"rest-api/config"
	"rest-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectToDb()

	api:= app.Group("/api")
	routes.ItemRoutes(api)

	app.Get("/",func(ctx *fiber.Ctx) error{
		return ctx.SendString("Server Running!")
	})

	app.Listen(":5000")
}