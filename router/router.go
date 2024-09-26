package router

import (
	"server/handler"
	// "server/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	
	api := app.Group("/api", logger.New())
	app.Use(cors.New())
	api.Get("/", handler.Hello)
	// api.Get("/mail", handler.SendMail)
	api.Post("/twoFA", handler.OtpHandler)
	api.Post("/validateOTP",handler.ValidationHandler)

	user := api.Group("/user")
	user.Post("/create", handler.CreateUser)
	user.Post("/login", handler.Login)
	user.Post("/profile", handler.GetUserProfile)
	user.Post("/updatePassword", handler.UpdatePassword)
	user.Post("/checkUsername", handler.CheckIfUsernameExists)
	user.Post("/checkGmail", handler.CheckIfGmailExists)
	user.Post("/checkGithub", handler.CheckIfGithubExists)
	user.Post("/recoverPassword", handler.PasswordRecovery)
	user.Post("/tempLogin", handler.TempLogin)
	user.Post("/updatePassword", handler.UpdatePassword)
	
	company := api.Group("/company")
	company.Get("/search", handler.SearchCompanies)
	company.Get("/compute/:companyID", handler.ComputeData)
	company.Get("/:companyID/financials", handler.FetchFinancialData)

	user.Get("/search-history", handler.ListSearchHistories)
	user.Get("/search-histories/:id", handler.GetSearchHistoryByID)
}
