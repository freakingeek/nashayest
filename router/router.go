package router

import (
	"github.com/sttatusx/nashayest/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api_v1.0", logger.New())

	// Routes
	api.Get("/", handler.CheckIllegalWords)

	api.Get("/words", handler.GetAllTheIllegalWords)
	api.Post("/words", handler.AddNewIllegalWord)
}
