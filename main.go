package main

import (
	"github.com/sttatusx/nashayest/router"

	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
