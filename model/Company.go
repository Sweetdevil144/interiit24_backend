package models

import (
	"time"
	"gorm.io/gorm"
)

type Company struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"type:varchar(255);not null;uniqueIndex"`
	Country     string         `gorm:"type:varchar(100);index"`
	CountryCode string         `gorm:"type:char(2);index"`
	MarketCap   float64        `gorm:"type:decimal(20,2)"`
	Diversity   float64        `gorm:"type:decimal(5,2)"`
	Financials  []FinancialData `gorm:"foreignKey:CompanyID;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
type FinancialData struct {
	ID          uint      `gorm:"primaryKey"`
	CompanyID   uint      `gorm:"not null;index"`
	Year        int       `gorm:"not null;index"`
	StockPrice  float64   `gorm:"type:decimal(10,2)"`
	Expense     float64   `gorm:"type:decimal(20,2)"`
	Revenue     float64   `gorm:"type:decimal(20,2)"`
	MarketShare float64   `gorm:"type:decimal(5,2)"`
	Company     Company   `gorm:"foreignKey:CompanyID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}