package handler

import (
	"server/database"
	"server/model"
	"github.com/gofiber/fiber/v2"
)

func SaveSearchHistory(userID uint, searchTerm string) error {
	db:=database.DB
	history := model.SearchHistory{UserID: userID, SearchTerm: searchTerm}
	return db.Create(&history).Error
}

func GetSearchHistory(userID uint) ([]model.SearchHistory, error) {
	db:=database.DB
	var histories []model.SearchHistory
	err := db.Model(&model.SearchHistory{}).
	Where("user_id = ?", userID).
	Order("timestamp desc").
	Find(&histories).Error
	return histories, err
}

func getUserIDFromContext(c *fiber.Ctx) uint {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return 0 
	}
	return userID
}

func GetUserSearchLog(c *fiber.Ctx) error {
	userID := getUserIDFromContext(c)
	logs, err := GetSearchHistory(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}
	return c.JSON(logs)
}

