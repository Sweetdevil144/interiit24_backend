package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Gmail    string `gorm:"uniqueIndex;not null" json:"gmail"`
	Github    string `gorm:"uniqueIndex" json:"github"`
	Password string `gorm:"not null" json:"password"`
	Name    string `json:"name"`
}
