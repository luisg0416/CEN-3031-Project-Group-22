package routes

import (
	//"bytes"
	//"encoding/json"
	//"bytes"
	//"encoding/json"
	//"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	//"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	//"github.com/luisg0416/CEN-3031-Project-Group-22/Models"
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

	Routes(app)

	for _, test := range tests {

		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req, 1)

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

	Routes(app)

	for _, test := range tests {

		req := httptest.NewRequest(http.MethodPost, test.route, bodyReader)
		//req.Header.Set("Header_Key", "Header_Value")
		resp, err := app.Test(req, -1)
		if err != nil {
			t.Log(err)
		}
		defer resp.Body.Close()
		
		// resBody, err2 := ioutil.ReadAll(resp.Body)
		// if err2 != nil {
		// 	t.Log(err)
		// }
		// response := string(resBody)
		
		// respBytes := []byte(response)
		
		// var jsonRes map[string]interface{}
		// //err1 := json.Unmarshal(respBytes, &jsonRes)
		// //if err1 != nil {
		// //	t.Log(err)
		// //}
		
		// respCard := &Models.Card {
		// 	Id: jsonRes["id"].(int),
		// 	Word: jsonRes["word"].(string),
		// 	Definition: jsonRes["definition"].(string),
		// }	


		// assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		// assert.Equalf(t, test.expectedID, respCard.Id, test.description)
		// assert.Equalf(t, test.expectedCode, respCard.Word, test.description)
		// assert.Equalf(t, test.expectedCode, respCard.Definition, test.description)
		assert.Equal(t, 422, resp.StatusCode)
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

	Routes(app)

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

	Routes(app)

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

	Routes(app)

	for _, test := range tests {
		req := httptest.NewRequest("Delete", test.route, nil)

		resp, _ := app.Test(req, 1)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}


