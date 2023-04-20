package controllers

import (
	//"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/gofiber/fiber/v2"
	//"github.com/luisg0416/CEN-3031-Project-Group-22/initializers"
	"github.com/luisg0416/CEN-3031-Project-Group-22/models"
)

type User struct {
	Name string `csv:"username" json:"username"`
	Password string `csv:"password" json:"password"`
}

func (u *UserCards) Login(c *fiber.Ctx) error {
	u.Cards = nil
	someuser := &User{}

	if err := c.BodyParser(someuser); err != nil {
		return err
	}
	
	data,err := os.ReadFile("storage/users.csv")
	if (err != nil) {
		panic(err)
	}
	var users []User
	_ = gocsv.UnmarshalBytes(data, &users)

	for _, user := range users {
		if (user.Name == someuser.Name) {
			if (user.Password == someuser.Password) {
				u.Username = someuser.Name
				u.InitializeCards()
				return c.SendStatus(200)
			} else {
				return c.SendStatus(403)	
			}
		}
	}
	return c.SendStatus(402)
}

func (u *UserCards) InitializeCards() {
	data,err := os.ReadFile("storage/flashCards.csv")
	if (err != nil) {
		panic(err)
	}
	var cardsWuser []models.CardWUser
	_ = gocsv.UnmarshalBytes(data, &cardsWuser)

	for _,card := range cardsWuser {
		if card.User == u.Username {
			u.Cards = append(u.Cards, models.Card{
				Id: card.Id,
				Word: card.Word,
				Definition: card.Definition,
			})
		}
	}
}

func (u *UserCards) CreateUser(c *fiber.Ctx) error {
	return nil
}

