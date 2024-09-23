package utils

import (
	"fmt"
	"gorm.io/gorm/clause"
	"server/database"
	"server/model"
)

func TwoFA(tempToken string) error {
	otp := GenerateOTP()
	err := SendMail(tempToken, otp)
	if err != nil {
		return err
	}
	return InsertOrUpdateOTP(tempToken, otp)
}

func ValidateAndDeleteOTP(tempToken, otp string) error {
	db := database.DB
	var deletedRows []model.OtpQueue
	db.Unscoped().Clauses(clause.Returning{}).Where("temp_token = ? AND otp = ?", tempToken, otp).Delete(&deletedRows)
	if len(deletedRows) == 0 {
		return fmt.Errorf("bad request")
	}
	return nil

}
