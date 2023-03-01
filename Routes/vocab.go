package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luisg0416/CEN-3031-Project-Group-22/Models"
)


func CreateFlashCard(c *fiber.Ctx, flashCards []Models.Card) error {
	flashCard := &Models.Card{}
		
	if err := c.BodyParser(flashCard); err != nil {
		return err
	}

	flashCard.Id = len(flashCards) + 1

	flashCards = append(flashCards, *flashCard)

	//fmt.Println(flashCard)

	return c.JSON(flashCards)
}
