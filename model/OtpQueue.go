package model

import "gorm.io/gorm"

// User struct
type OtpQueue struct {
	gorm.Model
	TempToken string `gorm:"uniqueIndex;not null" json:"temp_token"`
	Otp    string `gorm:"not null" json:"otp"`
}
