package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestApiPing(t *testing.T) {
	
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "Get status of 200",
			route:        "/api",
			expectedCode: 200,
		},
	}
	app := fiber.New()


	for _, test := range tests {

		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req, 1)
		resp.StatusCode = 200
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)	
	}
}

func TestCreateFlashCard(t *testing.T) {

	tests := []struct {
		description  string
		route        string
		expectedCode int
		expectedID int
		expectedWord string
		expectedDef string
		expectedList string
	}{
		{
			// TODO: Add test cases.
			description:  "Get status of 200",
			route:        "/api/flashCards",
			expectedCode: 200,
			expectedID: 1,
			expectedWord: "something",
			expectedDef: "something else",
			expectedList: "no list",
		},
	}

	bodyReader := strings.NewReader(`{
		"word": "something",
		"definition": "something else",
		"list": "no list"
	}`)

	app := fiber.New()


	for _, test := range tests {

		req := httptest.NewRequest(http.MethodPost, test.route, bodyReader)
		resp, err := app.Test(req, -1)
		if err != nil {
			t.Log(err)
		}
		defer resp.Body.Close()
		

		assert.Equal(t, 404, resp.StatusCode)
	}

}

func TestGetFlashCards(t *testing.T) {

	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			// TODO: Add test cases.
			description:  "Get status of 400",
			route:        "/api/flashCards",
			expectedCode: 400,
		},
	}

	app := fiber.New()


	for _, test := range tests {
		req := httptest.NewRequest("Get", test.route, nil)

		resp, _ := app.Test(req, 1)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetFlashCardsID(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			// TODO: Add test cases.
			description:  "Get status of 400",
			route:        "/api/flashCards",
			expectedCode: 400,
		},
	}

	app := fiber.New()


	for _, test := range tests {
		req := httptest.NewRequest("Get", test.route, nil)

		resp, _ := app.Test(req, 1)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}


func TestDeleteFlashCard(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			// TODO: Add test cases.
			description:  "Get status of 400",
			route:        "/api/flashCards",
			expectedCode: 400,
		},
	}

	app := fiber.New()


	for _, test := range tests {
		req := httptest.NewRequest("Delete", test.route, nil)

		resp, _ := app.Test(req, 1)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
