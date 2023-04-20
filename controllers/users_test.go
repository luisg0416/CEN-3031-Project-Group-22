package controllers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestUserCards_Login(t *testing.T) {
	tests := []struct {
		UserName    string
		Password string
	}{
		{
			UserName: "kory",
			Password: "12345",
		},{
			UserName: "mari",
			Password: "65432",
		},{
			UserName: "mathew",
			Password: "9876",
		},{
			UserName: "luis",
			Password: "78563",
		},
	}
	for _, tt := range tests {
		t.Run(tt.UserName, func(t *testing.T) {
		})
	}
}

func TestUserCards_InitializeCards(t *testing.T) {
	tests := []struct {
		name string
		u    *UserCards
	}{	
	}
	test := []struct {
		UserName    string
		Password string
	}{
		{
			UserName: "kory",
			Password: "12345",
		},{
			UserName: "mari",
			Password: "65432",
		},{
			UserName: "mathew",
			Password: "9876",
		},
	}
	if (len(test) ==0) {
		one := 0
		one++		
	}	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.u.InitializeCards()
		})
	}
}

func TestUserCards_CreateUser(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		u       *UserCards
		args    args
		wantErr bool
	}{		
	}
	test := []struct {
		UserName    string
		Password string
	}{
		{
			UserName: "kory1",
			Password: "12345",
		},{
			UserName: "mari4",
			Password: "65432",
		},{
			UserName: "mathew8",
			Password: "9876",
		},
	}
	if (len(test) ==0) {
		one := 0
		one++		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.CreateUser(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UserCards.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
