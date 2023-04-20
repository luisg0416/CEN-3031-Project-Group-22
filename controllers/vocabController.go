package controllers

import (
	"os"
	"github.com/gocarina/gocsv"
	"github.com/gofiber/fiber/v2"
	"github.com/luisg0416/CEN-3031-Project-Group-22/models"
)

var flashCards []models.Card

type UserCards struct {
	Cards []models.Card
	Username string
}

func ApiPing(c *fiber.Ctx) error {
	return c.Status(200).SendString("golang api up and running")
}


func (u *UserCards) CreateFlashCard(c *fiber.Ctx) error {
	if (u.Username == "") {
		return c.Status(400).SendString("No user")
	}
	
	flashCard := &models.Card{}
		
	if err := c.BodyParser(flashCard); err != nil {
		return err
	}

	flashCard.Id = len(u.Cards) + 1

	u.Cards = append(u.Cards, *flashCard)

	csvCards, err := os.Create("storage/flashCards.csv")
	if err != nil {
		panic(err)
	}
	defer csvCards.Close()

	newCard := models.CardWUser{
		Id: flashCard.Id,
		User: u.Username,
		Word: flashCard.Word,
		Definition: flashCard.Definition,
	}
	gocsv.MarshalFile(&newCard, csvCards)	


	//fmt.Println(flashCard)

	return c.SendStatus(200)
}

func (u *UserCards) GetFlashCards(c *fiber.Ctx) error {
	if len(u.Cards) == 0 {
		return c.Status(400).SendString("No FlashCards")		
	}

	return c.JSON(u.Cards)
}

func (u *UserCards) GetFlashCardsID(c *fiber.Ctx) error {
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

func (u *UserCards) DeleteFlashCard(c *fiber.Ctx) error {
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
