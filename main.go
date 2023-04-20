package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/luisg0416/CEN-3031-Project-Group-22/controllers"
	"github.com/luisg0416/CEN-3031-Project-Group-22/initializers"
	//"github.com/luisg0416/CEN-3031-Project-Group-22/models"
)

func init() {
	initializers.LoadEnvVars()
	//initializers.ConnectToDB()
}

func main() {
	// Load templates
	engine := html.New("./views", ".html")

	user := controllers.UserCards{
		Username: "kory",
	}
	user.InitializeCards()

	// Create app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure app
	app.Static("/", "./public")

	// Routing
	app.Post("/api/login", user.Login)
	app.Post("/api/signup", user.CreateUser)

	app.Get("/api", controllers.ApiPing)
	app.Post("/api/flashCards", user.CreateFlashCard)
	app.Get("/api/flashCards", user.GetFlashCards)
	app.Get("/api/flashCards/:id", user.GetFlashCardsID)
	app.Delete("/api/flashCards/:id", user.DeleteFlashCard)

	app.Get("/api/tasks", controllers.FetchTasks)
	app.Post("/api/tasks", controllers.CreateTask)
	app.Get("/api/tasks/:id", controllers.FetchTask)
	app.Delete("/api/tasks/:id", controllers.DeleteTask)

	frontendRoutes := []string{
		"/",
		"/about",
		"/sign-in",
		"/quizzes",
		"/quizzes/cop4600",
		"/quizzes/phy2053",
		"/flashcards",
		"/flashcards/cop4600",
		"/flashcards/phy2053",
	}

	for _, route := range frontendRoutes {
		app.Get(route, controllers.Home)
	}

	// Start app
	app.Listen(":" + os.Getenv("PORT"))

}
