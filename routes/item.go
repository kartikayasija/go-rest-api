package routes

import (
	"rest-api/controller"

	"github.com/gofiber/fiber/v2"
)

func ItemRoutes(app fiber.Router) {
	itemRoutes := app.Group("/items")
	itemRoutes.Get("/getAllItems",controller.GetAllItems)
	itemRoutes.Post("/postItem",controller.PostItem)
}