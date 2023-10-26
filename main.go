package main

import (
	"fiber/database"
	"fiber/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiber.Ctx) error {

	return c.SendString("Getting atrted wiht fiber")
}
// func setupRoutes( app *fiber.App){
// 	app.Get("/api",Welcome)
// 	// User endpoints
// 	app.Post("/api/users",routes.CretaeUser)
// }

func main() {
	database.ConnectDB()

	app := fiber.New()

	// setupRoutes(app)
	app.Get("/api", Welcome)
	app.Get("/api/getuser", routes.GetUsers)
	app.Post("/api/users",routes.CretaeUser)
	app.Get("/api/user/:id",routes.Getuser)

	log.Fatal(app.Listen(":4000"))
}
