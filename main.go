package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	allTheIllegalWords := []string{"عوضی", "بیشعور"} // Just some random words to test the result

	app.Get("/", func(c *fiber.Ctx) error {
		text := c.Query("text")
		illegalWords := []string{}

		if text != "" {
			words := strings.Split(text, " ")

			for _, word := range words {
				for _, illegal := range allTheIllegalWords {
					if illegal == word {
						illegalWords = append(illegalWords, word)
					}
				}
			}
		}

		return c.JSON(fiber.Map{
			"isIllegal":    len(illegalWords) > 0,
			"illegalWords": illegalWords,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
