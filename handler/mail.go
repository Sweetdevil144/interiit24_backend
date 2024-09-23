package handler

import (
	"fmt"
	// "github.com/gofiber/fiber/v2"
	smtp "net/smtp"
	"server/config"
)

func SendOTP(to, otp string) error {
	password := config.Config("SMTP_PASSWORD")
	from := config.Config("SMTP_EMAIL")
	smtpHost := config.Config("SMTP_HOST")
	smtpPort:=config.Config("SMTP_PORT")
	message := []byte(fmt.Sprintf("Subject: OTP\r\n\r\notp : %s\r\n", otp))
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		return err
	}
	fmt.Println("Email Sent Successfully")
	return nil
}

func SendMail(tempToken,otp string) error {
	_,gmail,err:=DeserialiseTempToken(tempToken)
	if err!=nil{
		return err
	}
	return SendOTP(gmail, otp)
}

// func SendMail(c *fiber.Ctx) error {
// 	err := SendOTP("devgup04@gmail.com", "1234")
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{"message": err})
// 	}
// 	return c.Status(200).JSON(fiber.Map{"message": "success"})
// }
