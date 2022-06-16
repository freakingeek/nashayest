package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CheckIllegalWords(c *fiber.Ctx) error {
	text := c.Query("text")
	illegalWords := []string{}
	allTheIllegalWords := []string{"عوضی", "بیشعور"} // Just some random words to test the result

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
}
