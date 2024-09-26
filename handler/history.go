package handler

import (
	"errors"
	"server/database"
	"server/model"
	"strconv"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)

func ListSearchHistories(c *fiber.Ctx) error {
	userIDStr := c.Params("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	var searchHistories []model.SearchHistory
	if err := database.DB.Where("user_id = ?", userID).
		Order("timestamp DESC").
		Limit(10).
		Find(&searchHistories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to fetch search histories"})
	}
	return c.JSON(searchHistories)
}
func GetSearchHistoryByID(c *fiber.Ctx) error {
	historyID := c.Params("id")
	userID := c.Locals("userID").(uint)
	var searchHistory model.SearchHistory
	if err := database.DB.Where("id = ? AND user_id = ?", historyID, userID).First(&searchHistory).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Search history not found"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to fetch search history"})
	}
	return c.JSON(searchHistory)
}
