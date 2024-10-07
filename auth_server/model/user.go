package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Gmail    string `gorm:"uniqueIndex;not null" json:"gmail"`
	Github    string `json:"github"`
	Password string `json:"password"`
	Name    string `gorm:"not null" json:"name"`
}
