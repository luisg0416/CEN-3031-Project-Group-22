//   http://127.0.0.1:3000

package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	//"github.com/luisg0416/CEN-3031-Project-Group-22/database"
	"github.com/luisg0416/CEN-3031-Project-Group-22/Routes"
)


func main() {

	fmt.Println("this is just so vscode shuts up about my imports")
	
	//database.ConnectDb()
	app := fiber.New()

	routes.Routes(app)

	log.Fatal(app.Listen(":3000"))
}