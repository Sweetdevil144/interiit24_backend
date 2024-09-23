package handler

import (
	"gorm.io/gorm"
	"server/model"
)

func SaveSearchHistory(userID uint, searchTerm string) error {
	history := models.SearchHistory{UserID: userID, SearchTerm: searchTerm}
	return DB.Create(&history).Error
}

func GetSearchHistory(userID uint) ([]models.SearchHistory, error) {
	var histories []models.SearchHistory
	err := DB.Model(&models.SearchHistory{}).
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
	logs, err := database.GetSearchHistory(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}
	return c.JSON(logs)
}

