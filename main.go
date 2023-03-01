//   http://127.0.0.1:3000

package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/luisg0416/CEN-3031-Project-Group-22/database"
	"github.com/luisg0416/CEN-3031-Project-Group-22/Models"
	//"github.com/luisg0416/CEN-3031-Project-Group-22/Routes"
)


var flashCards []Models.Card

func main() {

	fmt.Println("this is just so vscode shuts up about my imports")
	
	database.ConnectDb()
	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/api/flashCards", func(c *fiber.Ctx) error {
		flashCard := &Models.Card{}
	
		if err := c.BodyParser(flashCard); err != nil {
			return err
		}
	
		flashCard.Id = len(flashCards) + 1
	
		flashCards = append(flashCards, *flashCard)
	
		//fmt.Println(flashCard)
	
		return c.JSON(flashCards)
	})

	app.Get("/api/flashCards", func(c *fiber.Ctx) error {
		return c.JSON(flashCards)
	})

	app.Get("/api/flashCards/:id", func(c *fiber.Ctx) error {
		
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		var card Models.Card
		for _, card := range flashCards {
			if card.Id == id {
				return c.JSON(card)
			}
		}

		return c.JSON(card)
	})

	log.Fatal(app.Listen(":3000"))
}
