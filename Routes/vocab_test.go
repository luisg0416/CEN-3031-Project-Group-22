package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/luisg0416/CEN-3031-Project-Group-22/Models"
	"github.com/stretchr/testify/assert"
)

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

	card := &Models.Card {
		Id: 1, 
		Word: "something", 
		Definition: "something else", 
		List: "no list",
	}

	var jsonCard bytes.Buffer
	err := json.NewEncoder(&jsonCard).Encode(card)
	if err != nil {
		t.Fatal(err)
	}

	app := fiber.New()

	Routes(app)

	for _, test := range tests {
		
		req := httptest.NewRequest("Post", test.route, &jsonCard)
		req.Header.Set("Header_Key", "Header_Value")
		resp, _ := app.Test(req, 1)
		defer resp.Body.Close()
		
		resBody, _ := ioutil.ReadAll(resp.Body)
		response := string(resBody)
		
		respBytes := []byte(response)
		
		var jsonRes map[string]interface{}
		err := json.Unmarshal(respBytes, &jsonRes)
		if err != nil {
			fmt.Println(err)
		}
		
		respCard := &Models.Card{
			Id: jsonRes["id"].(int),
			Word: jsonRes["word"].(string),
			Definition: jsonRes["definition"].(string), 
			List: jsonRes["list"].(string),
		}	


		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		assert.Equalf(t, test.expectedID, respCard.Id, test.description)
		assert.Equalf(t, test.expectedCode, respCard.Word, test.description)
		assert.Equalf(t, test.expectedCode, respCard.Definition, test.description)
		assert.Equalf(t, test.expectedCode, respCard.List, test.description)
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
			description:  "Get status of 200",
			route:        "/api/flashCards",
			expectedCode: 200,
		},
	}

	app := fiber.New()

	Routes(app)

	for _, test := range tests {
		req := httptest.NewRequest("Get", test.route, nil)

		resp, _ := app.Test(req, 1)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
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
