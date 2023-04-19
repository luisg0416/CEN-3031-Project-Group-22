package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/luisg0416/CEN-3031-Project-Group-22/controllers"
	"github.com/luisg0416/CEN-3031-Project-Group-22/initializers"
)

func init() {
	initializers.LoadEnvVars()
	//initializers.ConnectToDB()
}

func main() {
	// Load templates
	engine := html.New("./views", ".html")

	// Create app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure app
	app.Static("/", "./public")

	// Routing
	app.Get("/api", controllers.ApiPing)
	app.Post("/api/flashCards", controllers.CreateFlashCard)
	app.Get("/api/flashCards", controllers.GetFlashCards)
	app.Get("/api/flashCards/:id", controllers.GetFlashCardsID)
	app.Delete("/api/flashCards/:id", controllers.DeleteFlashCard)
	
	app.Get("/api/tasks", controllers.FetchTasks)
	app.Post("/api/tasks", controllers.CreateTask)
	app.Get("/api/tasks/:id", controllers.FetchTask)
	app.Delete("/api/tasks/:id", controllers.DeleteTask)

	frontendRoutes := []string{
		"/",
		"/about",
		"/sign-in",
		"/sign-in/create-acc",
		"/flashcards",
	}

	for _, route := range frontendRoutes {
		app.Get(route, controllers.Home)
	}

	// Start app
	app.Listen(":" + os.Getenv("PORT"))

}

