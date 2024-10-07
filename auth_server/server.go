package main

import (
	"auth_server/cache"
	"auth_server/database"
	"auth_server/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()
	cache.Init()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":6969"))
}
