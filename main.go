package main

import (
	"github.com/sttatusx/nashayest/router"

	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}
