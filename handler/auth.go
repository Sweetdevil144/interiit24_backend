package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"server/config"
	"golang.org/x/crypto/bcrypt"
	// "fmt"
)

var serialKey = []byte(config.Config("JWT_SECRET"))
var gmailKey = []byte(config.Config("GMAIL_SECRET"))
var githubKey = []byte(config.Config("GITHUB_SECRET"))

func SerialiseUser(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	signedToken, err := token.SignedString(serialKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func DeserialiseUser(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(serialKey), nil
	})

	if err != nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims["username"].(string), nil
}

func CheckPasswordWithHash(storedPasswwordHash string, password string) bool{
	return bcrypt.CompareHashAndPassword([]byte(storedPasswwordHash),[]byte(password))==nil
}

func DeserialiseGmailToken(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(gmailKey), nil
	})

	if err != nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims["gmail"].(string), nil
}

func DeserialiseGithubToken(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(githubKey), nil
	})

	if err != nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims["github"].(string), nil
}