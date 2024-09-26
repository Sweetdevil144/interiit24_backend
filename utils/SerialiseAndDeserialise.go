package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"server/config"
)

var serialKey = []byte(config.Config("JWT_SECRET"))
var gmailKey = []byte(config.Config("GMAIL_SECRET"))
var githubKey = []byte(config.Config("GITHUB_SECRET"))
var tempTokenKey = []byte(config.Config("TEMPTOKEN_SECRET"))

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

func SerialiseGmailToken(gmail string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"gmail": gmail,
	})
	signedToken, err := token.SignedString(gmailKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
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

func SerialiseGithubToken(github string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"github": github,
	})
	signedToken, err := token.SignedString(githubKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
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

func SerialiseTempToken(username, gmail string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"gmail":    gmail,
	})
	signedToken, err := token.SignedString(tempTokenKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func DeserialiseTempToken(signedToken string) (string, string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(tempTokenKey), nil
	})

	if err != nil {
		return "", "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims["username"].(string), claims["gmail"].(string), nil
}
