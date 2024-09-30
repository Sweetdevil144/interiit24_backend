package handler

import (
	// "fmt"
	"auth_server/utils"

	"github.com/gofiber/fiber/v2"
)

func ValidationHandler(c *fiber.Ctx) error {
	var body struct {
		TempToken string `json:"temp_token"`
		Otp       string `json:"otp"`
	}
	c.BodyParser(&body)
	// fmt.Println(body.TempToken,body.Otp,body.OtpType)
	err := utils.ValidateAndDeleteOTP(body.TempToken, body.Otp, "login")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	username, _, err := utils.DeserialiseTempToken(body.TempToken)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := utils.SerialiseUser(username)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"token": token})
}

func OtpHandler(c *fiber.Ctx) error {
	var body struct {
		TempToken string `json:"temp_token"`
	}
	c.BodyParser(&body)
	err := utils.TwoFA(body.TempToken, "login")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}
	return c.Status(200).JSON(fiber.Map{"message": "otp generated successfully"})

}
