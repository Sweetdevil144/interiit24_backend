package config

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
)

func Config(key string) string {
	path,_:=os.Getwd()
	dir, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		fmt.Print(err)
	}
	environmentPath := filepath.Join(dir, ".env")
	err = godotenv.Load(environmentPath)
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
