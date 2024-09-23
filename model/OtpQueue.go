package model

import (
	"time"

	"gorm.io/gorm"
)

// User struct
type OtpQueue struct {
	gorm.Model
	TempToken string `gorm:"uniqueIndex;not null" json:"temp_token"`
	Otp    string `gorm:"not null" json:"otp"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
}
