package database

import (
	"auth_server/config"
	"auth_server/model"
	"fmt"
	"strconv"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	p := config.Config("DB_PORT")
	port, err := strconv.Atoi(p)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database port: %v", err)
	}

	host := config.Config("DB_HOST")
	if host == "" {
		log.Fatal("DB_HOST environment variable is missing")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&model.User{}, &model.OtpQueue{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	fmt.Println("Connection Opened and Database Migrated")
	return db, nil
}
