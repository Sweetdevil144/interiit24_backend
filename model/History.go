package models

import (
	"time"
	"gorm.io/gorm"
)

type SearchHistory struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `gorm:"not null;index"`
	SearchTerm string        `gorm:"type:text;not null"`
	Timestamp time.Time      `gorm:"autoCreateTime"`
}

