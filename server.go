package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
	"server/cache"
	"server/database"
	"server/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "http://localhost:3000",
			AllowMethods: "GET,POST,PUT,DELETE",
		}))
	database.ConnectDB()
	cache.Init()
	router.SetupRoutes(app)
	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}
	log.Fatal(app.Listen(":" + port))
}
