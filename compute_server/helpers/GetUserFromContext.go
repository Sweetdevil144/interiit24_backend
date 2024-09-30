package helpers

import (
	"compute_server/database"
	"compute_server/model"
	"compute_server/utils"

	"github.com/gofiber/fiber/v2"
)

func GetUserFromContext(c *fiber.Ctx) (uint, error) {
	token := c.Get("Authorization")[7:]
	username, err := utils.DeserialiseUser(token)
	if err != nil {
		return 0, err
	}
	var user model.User
	if err := database.DB.Select("id").Where("username = ?", username).First(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}
