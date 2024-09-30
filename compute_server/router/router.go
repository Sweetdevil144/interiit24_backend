package router

import (
	"compute_server/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())
	app.Use(cors.New())
	api.Get("/", handler.Hello)

	user := api.Group("/user")

	company := api.Group("/company")
	company.Get("/search", handler.SearchCompanies)
	company.Get("/compute/:companyID", handler.ComputeData)
	company.Get("/:companyID/financials", handler.FetchFinancialData)

	user.Get("/search-history", handler.ListSearchHistories)
	user.Get("/search-histories/:id", handler.GetSearchHistoryByID)
}
