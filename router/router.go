package router

import (
	"server/handler"
	// "server/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	user := api.Group("/user")
	user.Post("/create", handler.CreateUser)
	user.Post("/login", handler.Login)
	user.Post("/updatePassword",handler.UpdatePassword)
}
