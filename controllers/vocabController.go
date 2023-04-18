package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luisg0416/CEN-3031-Project-Group-22/models"
)

var flashCards []models.Card

func ApiPing(c *fiber.Ctx) error {
	return c.Status(200).SendString("golang api up and running")
}


func CreateFlashCard(c *fiber.Ctx) error {
	flashCard := &models.Card{}
		
	if err := c.BodyParser(flashCard); err != nil {
		return err
	}

	flashCard.Id = len(flashCards) + 1

	flashCards = append(flashCards, *flashCard)

	//fmt.Println(flashCard)

	return c.SendStatus(200)
}

func GetFlashCards(c *fiber.Ctx) error {
	if len(flashCards) == 0 {
		return c.Status(400).SendString("No FlashCards")
	}

	return c.JSON(flashCards)
}

func GetFlashCardsID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(401).SendString("Invalid id")
	}

	var card models.Card
	for _, card := range flashCards {
		if card.Id == id {
			return c.JSON(card)
		}
	}

	if card.Id == 0 {
		return c.Status(401).SendString("ID does not exist")
	}

	return c.JSON(card)
} 

func DeleteFlashCard(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(401).SendString("Invalid Id")
	}
	var cards []models.Card
	found := false
	for _, card := range flashCards {
		if card.Id == id {
			found = true
		} else {
			cards = append(cards, card)
		}
	}
	if (!found) {
		return c.Status(401).SendString("ID does not exist")
	}

	flashCards = cards
	return c.SendStatus(200)
}
