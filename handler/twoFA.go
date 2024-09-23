package handler

import (
	"fmt"
	"math/rand"
	"server/database"
	"server/model"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GenerateOTP() string {
	otp := ""
	for i := 0; i < 4; i++ {
		otp += strconv.Itoa(rand.Intn(10))
	}
	return otp
}

func InsertOrUpdateOTP(tempToken, otp string) error {
	db := database.DB
	err := db.Where(model.OtpQueue{TempToken: tempToken}).
		Assign(model.OtpQueue{TempToken: tempToken, Otp: otp, ExpiresAt: time.Now().Add(10 * time.Minute)}).
		FirstOrCreate(&model.OtpQueue{}).Error
	return err
}

func TwoFA(tempToken string) error {
	return InsertOrUpdateOTP(tempToken, GenerateOTP())
}

func ValidateAndDeleteOTP(tempToken, otp string) error {
	db := database.DB
	var deletedRows []model.OtpQueue
	db.Unscoped().Clauses(clause.Returning{}).Where("temp_token = ? AND otp = ?", tempToken,otp).Delete(&deletedRows)
	if len(deletedRows) == 0 {
		return fmt.Errorf("bad request")
	}
	return nil

}

func ValidationHandler(c *fiber.Ctx) error {
	var body struct {
		TempToken string `json:"temp_token"`
		Otp       string `json:"otp"`
	}
	c.BodyParser(&body)
	err := ValidateAndDeleteOTP(body.TempToken, body.Otp)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "otp validated"})
}

func OtpHandler(c *fiber.Ctx) error {
	var body struct {
		TempToken string `json:"temp_token"`
	}
	c.BodyParser(&body)
	err := TwoFA(body.TempToken)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}
	return c.Status(200).JSON(fiber.Map{"message": "otp generated successfully"})

}
