package routes

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("golang api up and running")
	})

	app.Post("/api/flashCards", CreateFlashCard)

	app.Get("/api/flashCards", GetFlashCards)

	app.Get("/api/flashCards/:id", GetFlashCardsID)
}