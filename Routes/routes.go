package routes

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	app.Get("/api", ApiPing)

	app.Post("/api/flashCards", CreateFlashCard)

	app.Get("/api/flashCards", GetFlashCards)

	app.Get("/api/flashCards/:id", GetFlashCardsID)
}