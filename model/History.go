package model

import (
	"time"
	// "gorm.io/gorm"
)

type SearchHistory struct {
	ID        uint      `gorm:"primaryKey"`
	UserID       uint                   `gorm:"not null;index"` 
	User         User                   `gorm:"foreignKey:UserID"`
	CompanyID    uint                 `gorm:"type:text;not null"`
	Company      Company                `gorm:"foreignKey:CompanyID"`
	Timestamp time.Time `gorm:"autoCreateTime"`
	StoredResult map[string]interface{} `gorm:"type:jsonb"` 
}
