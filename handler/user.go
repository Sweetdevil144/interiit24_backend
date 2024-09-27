package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"server/database"
	"server/model"
	"server/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type userInfo struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Gmail    string `json:"gmail"`
	Github   string `json:"github"`
}

type updatePasswordInfo struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
type passwordRecoveryInfo struct {
	Otp         string `json:"otp"`
	TempToken   string `json:"temp_token"`
	NewPassword string `json:"new_password"`
}

type loginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type tempLoginInfo struct {
	Username string `json:"username"`
	Gmail    string `json:"gmail"`
}

type GitHubAccessTokenResponse struct {
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
    Scope       string `json:"scope"`
}

type GitHubUserProfile struct {
    Login string `json:"login"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func GitHubLogin(c *fiber.Ctx) error {
	type request struct {
		Code string `json:"code"`
	}

	var data request
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	tokenResponse, err := exchangeCodeForAccessToken(data.Code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to exchange code for token"})
	}

	userData, err := fetchGitHubUserData(tokenResponse.AccessToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch GitHub user data"})
	}

	return c.JSON(userData)
}

func exchangeCodeForAccessToken(code string) (*GitHubAccessTokenResponse, error) {
    clientID := []byte(config.Config(("GITHUB_CLIENT_ID")))
    clientSecret := []byte(config.Config(("GITHUB_CLIENT_SECRET")))
    url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Add("Accept", "application/json")
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()
    var tokenResponse GitHubAccessTokenResponse
    if err := json.NewDecoder(res.Body).Decode(&tokenResponse); err != nil {
        return nil, err
    }
    return &tokenResponse, nil
}

func fetchGitHubUserData(accessToken string) (*GitHubUserProfile, error) {
    url := "https://api.github.com/user"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Add("Authorization", "Bearer "+accessToken)
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()
    var userProfile GitHubUserProfile
    if err := json.NewDecoder(res.Body).Decode(&userProfile); err != nil {
        return nil, err
    }
    return &userProfile, nil
}

func GetUserProfile(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}
	c.BodyParser(&body)
	username, err := utils.DeserialiseUser(body.Token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid token not found"})
	}
	db := database.DB
	var result model.User
	queryRes := db.Select("username", "name", "gmail", "github").First(&result, &model.User{Username: username})
	if queryRes.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "user not found"})
	}
	return c.Status(200).JSON(fiber.Map{
		"username": result.Username,
		"name":    result.Name,
		"gmail":    result.Gmail,
		"github":    result.Github,
	})
}

func CreateUser(c *fiber.Ctx) error {
	body := new(userInfo)
	c.BodyParser(&body)
	token, _ := utils.SerialiseUser(body.Username)
	db := database.DB
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 4)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "error",
		})
	}
	body.Password = string(hashedPassword[:])
	gmail,_:=utils.DeserialiseGmailToken(body.Gmail)
	newUser := model.User{
		Name:     body.Name,
		Gmail:    gmail,
		Password: body.Password,
		Username: body.Username,
	}
	if body.Github != "" {
		// gmail,_:=utils.DeserialiseGithubToken(body.Gmail)
		newUser.Github = body.Github
	}

	queryRes := db.Create(&newUser)

	if queryRes.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": queryRes.Error,
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"token": token,
	})
}

func LoginWithPassword(c *fiber.Ctx) error {
	var body loginInfo
	c.BodyParser(&body)
	db := database.DB
	var result model.User
	// fmt.Println("username : ",body.Username)
	queryRes := db.First(&result, &model.User{Username: body.Username})
	// fmt.Println(queryRes.RowsAffected)
	// fmt.Println(result)
	if queryRes.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "user not found"})
	}
	if utils.CheckPasswordWithHash(result.Password, body.Password) {
		tempToken, _ := utils.SerialiseTempToken(result.Username, result.Gmail)
		err := utils.TwoFA(tempToken, "login")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "couldnt issue 2FA"})
		}
		return c.Status(200).JSON(fiber.Map{
			"temp_token": tempToken,
			"expires_at": time.Now().Add(10 * time.Minute),
		})
	} else {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid password",
		})
	}

}
func LoginWithGmail(c *fiber.Ctx) error {
	var body loginInfo
	c.BodyParser(&body)
	gmail, err := utils.DeserialiseGmailToken(body.Token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	db := database.DB
	var result model.User
	queryRes := db.First(&result, &model.User{Gmail: gmail})

	if queryRes.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "user not found"})
	} else {
		tempToken, _ := utils.SerialiseTempToken(result.Username, result.Gmail)
		err := utils.TwoFA(tempToken, "login")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "couldnt issue 2FA"})
		}
		return c.Status(200).JSON(fiber.Map{
			"temp_token": tempToken,
			"expires_at": time.Now().Add(10 * time.Minute),
		})
	}
}
func LoginWithGithub(c *fiber.Ctx) error {
	var body loginInfo
	c.BodyParser(&body)
	github, err := utils.DeserialiseGithubToken(body.Token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	db := database.DB
	var result model.User
	queryRes := db.First(&result, &model.User{Github: github})

	if queryRes.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "user not found"})
	} else {
		tempToken, _ := utils.SerialiseTempToken(result.Username, result.Gmail)
		err := utils.TwoFA(tempToken, "login")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "couldnt issue 2FA"})
		}
		return c.Status(200).JSON(fiber.Map{
			"temp_token": tempToken,
			"expires_at": time.Now().Add(10 * time.Minute),
		})
	}
}

func Login(c *fiber.Ctx) error {
	loginMethod := c.Get("Login-Method")
	if loginMethod == "password" {
		return LoginWithPassword(c)
	} else if loginMethod == "gmail" {
		return LoginWithGmail(c)
	} else if loginMethod == "github" {
		return LoginWithGithub(c)
	} else {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid login method"})
	}
}

func CheckIfUsernameExists(c *fiber.Ctx) error {
	var body userInfo
	c.BodyParser(&body)
	db := database.DB
	var result model.User
	queryRes := db.First(&result, &model.User{Username: body.Username})
	return c.Status(200).JSON(fiber.Map{"userExists": queryRes.RowsAffected == 1})
}
func CheckIfGmailExists(c *fiber.Ctx) error {
	var body userInfo
	c.BodyParser(&body)
	db := database.DB
	var result model.User
	queryRes := db.First(&result, &model.User{Gmail: body.Gmail})
	fmt.Println(result)
	return c.Status(200).JSON(fiber.Map{"userExists": queryRes.RowsAffected == 1})
}
func CheckIfGithubExists(c *fiber.Ctx) error {
	var body userInfo
	c.BodyParser(&body)
	db := database.DB
	var result model.User
	queryRes := db.First(&result, &model.User{Github: body.Github})
	return c.Status(200).JSON(fiber.Map{"userExists": queryRes.RowsAffected == 1})
}

func PasswordRecovery(c *fiber.Ctx) error {
	var body passwordRecoveryInfo
	c.BodyParser(&body)
	db := database.DB
	var result model.User
	username, _, err := utils.DeserialiseTempToken(body.TempToken)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid temp_token"})
	}
	err = utils.ValidateAndDeleteOTP(body.TempToken, body.Otp, "recovery")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid temp_token/otp"})

	}
	queryRes := db.First(&result, &model.User{Username: username})
	if queryRes.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "user doesnt exist"})
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 4)
	result.Password = string(hashedPassword[:])
	queryRes = db.Save(&result)
	if queryRes.RowsAffected == 0 {
		return c.Status(502).JSON(fiber.Map{"message": "couldnt update password"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "password updated successfully"})
}

func TempLogin(c *fiber.Ctx) error {
	var body tempLoginInfo
	c.BodyParser(&body)
	db := database.DB
	var result model.User
	queryRes := db.First(&result, &model.User{Username: body.Username, Gmail: body.Gmail})
	if queryRes.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "invalid username or email"})
	}
	tempToken, _ := utils.SerialiseTempToken(result.Username, result.Gmail)
	err := utils.TwoFA(tempToken, "recovery")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "couldnt issue 2FA"})
	}
	return c.Status(200).JSON(fiber.Map{
		"temp_token": tempToken,
		"expires_at": time.Now().Add(10 * time.Minute),
	})
}

func UpdatePassword(c *fiber.Ctx) error {
	var body updatePasswordInfo
	c.BodyParser(&body)
	authHeader := c.Get("Authorization")[7:]
	username, err := utils.DeserialiseUser(authHeader)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid token"})
	}
	db := database.DB
	var result model.User
	queryRes := db.First(&result, &model.User{Username: username})
	if queryRes.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "invalid username"})
	}
	if !utils.CheckPasswordWithHash(result.Password, body.OldPassword) {
		return c.Status(400).JSON(fiber.Map{"message": "invalid password"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 4)
	result.Password = string(hashedPassword[:])
	queryRes = db.Save(&result)
	if queryRes.RowsAffected == 0 {
		return c.Status(502).JSON(fiber.Map{"message": "couldnt update password"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "password updated successfully"})
}
