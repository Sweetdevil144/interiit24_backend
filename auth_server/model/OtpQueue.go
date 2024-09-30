package model

import (
	"time"
	"gorm.io/gorm"
)

type OtpQueue struct {
	gorm.Model
	TempToken string    `gorm:"not null;primaryKey" json:"temp_token"`
	Otp       string    `gorm:"not null" json:"otp"`
	OtpType   string    `gorm:"not null;primaryKey" json:"otp_type"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
}
