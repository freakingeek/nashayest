package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sttatusx/nashayest/database"
	"github.com/sttatusx/nashayest/models"
)

func CheckIllegalWords(c *fiber.Ctx) error {
	text := c.Query("text")
	illegalWords := []string{}

	// Get the illegal words from database
	allTheIllegalWords := []models.Word{}
	database.Database.Find(&allTheIllegalWords)

	if text != "" {
		words := strings.Split(text, " ")

		for _, word := range words {
			for _, illegal := range allTheIllegalWords {
				if illegal.Title == word {
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

func AddNewIllegalWord(c *fiber.Ctx) error {
	var word models.Word

	if err := c.BodyParser(&word); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	database.Database.Create(&word)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ID":    word.ID,
		"title": word.Title,
	})
}

func GetAllTheIllegalWords(c *fiber.Ctx) error {
	words := []models.Word{}
	database.Database.Find(&words)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": len(words),
		"words": words,
	})
}

func GetAnIllegalWordById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "You should provide an id with type of integer",
		})
	}

	var word models.Word

	database.Database.Find(&word, "id = ?", id)

	if word.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Word not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":    word.ID,
		"title": word.Title,
	})
}

func DeleteAnIllegalWordById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "You should provide an id with type of integer",
		})
	}

	var word models.Word

	database.Database.Find(&word, "id = ?", id)

	if word.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Word not found",
		})
	}

	if err = database.Database.Delete(&word).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":    word.ID,
		"title": word.Title,
	})
}
