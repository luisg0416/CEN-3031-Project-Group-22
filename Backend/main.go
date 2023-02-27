package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type card struct {
	id         int    `json:"id"`
	word       string `json:"word"`
	definition string `json:"definition"`
}

func main() {

	app := fiber.New()

	flashCards := []card{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/api/flashCards", func(c *fiber.Ctx) error {
		flashCard := &card{}

		if err := c.BodyParser(flashCard); err != nil {
			return err
		}

		flashCard.id = len(flashCards) + 1

		flashCards = append(flashCards, *flashCard)

		return c.JSON(flashCard)
	})

	log.Fatal(app.Listen(":3000"))
}
