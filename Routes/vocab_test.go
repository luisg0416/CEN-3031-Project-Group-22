package routes

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestCreateFlashCard(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateFlashCard(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateFlashCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFlashCards(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetFlashCards(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetFlashCards() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFlashCardsID(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetFlashCardsID(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetFlashCardsID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
