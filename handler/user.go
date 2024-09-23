package handler

import (
	// "strconv"
	_ "fmt"
	"server/database"
	"server/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type userInfo struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Gmail    string `json:"gmail"`
	Github   string `json:"github"`
}

type updatePasswordInfo struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type loginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func CreateUser(c *fiber.Ctx) error {
	body := new(userInfo)
	c.BodyParser(&body)
	token, _ := SerialiseUser(body.Username)
	db := database.DB
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 4)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "error",
		})
	}
	body.Password = string(hashedPassword[:])

	var newUser model.User
	newUser.Name = body.Name
	newUser.Gmail = body.Gmail
	if body.Github != "" {
		newUser.Github = body.Github
	}
	newUser.Password = body.Password
	newUser.Username = body.Username

	queryRes := db.Create(&newUser)

	if queryRes.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": queryRes.Error,
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"token": token,
	})
}

func LoginWithPassword(c *fiber.Ctx) error {
	var body loginInfo
	c.BodyParser(&body)

	db := database.DB
	var result model.User
	queryRes := db.First(&result, &model.User{Username: body.Username})
	if queryRes.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "body not found"})
	}
	if CheckPasswordWithHash(result.Password, body.Password) {
		token, _ := SerialiseUser(body.Username)
		return c.Status(200).JSON(fiber.Map{
			"token": token,
		})
	} else {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid password",
		})
	}

}
func LoginWithGmail(c *fiber.Ctx) error {
	var body loginInfo
	c.BodyParser(&body)
	gmail, err := DeserialiseGmailToken(body.Token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	db := database.DB
	var result model.User
	queryRes := db.First(&result, &model.User{Gmail: gmail})

	if queryRes.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "user not found"})
	} else {
		token, _ := SerialiseUser(result.Username)
		return c.Status(200).JSON(fiber.Map{
			"token": token,
		})
	}
}
func LoginWithGithub(c *fiber.Ctx) error {
	var body loginInfo
	c.BodyParser(&body)
	github, err := DeserialiseGithubToken(body.Token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	db := database.DB
	var result model.User
	queryRes := db.First(&result, &model.User{Github: github})

	if queryRes.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "user not found"})
	} else {
		token, _ := SerialiseUser(result.Username)
		return c.Status(200).JSON(fiber.Map{
			"token": token,
		})
	}
}

func Login(c *fiber.Ctx) error {
	loginMethod := c.Get("Login-Method")
	if loginMethod == "password" {
		return LoginWithPassword(c)
	} else if loginMethod == "gmail" {
		return LoginWithGmail(c)
	} else if loginMethod == "github" {
		return LoginWithGithub(c)
	} else {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid login method"})
	}
}

func CheckIfUsernameExists(c *fiber.Ctx) error {
	var body userInfo
	c.BodyParser(&body)
	db := database.DB
	var result model.User
	queryRes := db.First(&result, &model.User{Username: body.Username})
	return c.Status(200).JSON(fiber.Map{"userExists": queryRes.RowsAffected == 1})
}
func CheckIfGmailExists(c *fiber.Ctx) error {
	var body userInfo
	c.BodyParser(&body)
	db := database.DB
	var result model.User
	queryRes := db.First(&result, &model.User{Gmail: body.Gmail})
	return c.Status(200).JSON(fiber.Map{"userExists": queryRes.RowsAffected == 1})
}
func CheckIfGithubExists(c *fiber.Ctx) error {
	var body userInfo
	c.BodyParser(&body)
	db := database.DB
	var result model.User
	queryRes := db.First(&result, &model.User{Github: body.Github})
	return c.Status(200).JSON(fiber.Map{"userExists": queryRes.RowsAffected == 1})
}

func PasswordRecovery(c *fiber.Ctx) error {
	var body userInfo
	c.BodyParser(&body)
	db := database.DB
	var result model.User
	gmail, err := DeserialiseGmailToken(body.Gmail)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid gmail token"})
	}
	queryRes := db.First(&result, &model.User{Username: body.Username, Gmail: gmail})
	if queryRes.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "invalid username or email"})
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 4)
	result.Password = string(hashedPassword[:])
	queryRes = db.Save(&result)
	if queryRes.RowsAffected == 0 {
		return c.Status(502).JSON(fiber.Map{"message": "couldnt update password"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "password updated successfully"})
}

func UpdatePassword(c *fiber.Ctx) error {
	var body updatePasswordInfo
	c.BodyParser(&body)
	authHeader := c.Get("Authorization")[7:]
	username, err := DeserialiseUser(authHeader)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid token"})
	}
	db := database.DB
	var result model.User
	queryRes := db.First(&result, &model.User{Username: username})
	if queryRes.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "invalid username"})
	}
	if !CheckPasswordWithHash(result.Password, body.OldPassword) {
		return c.Status(400).JSON(fiber.Map{"message": "invalid password"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 4)
	result.Password = string(hashedPassword[:])
	queryRes = db.Save(&result)
	if queryRes.RowsAffected == 0 {
		return c.Status(502).JSON(fiber.Map{"message": "couldnt update password"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "password updated successfully"})
}
