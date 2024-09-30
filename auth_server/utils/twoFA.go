package utils

import (
	"auth_server/database"
	"auth_server/model"
	"fmt"

	"gorm.io/gorm/clause"
)

func TwoFA(tempToken, otp_type string) error {
	otp := GenerateOTP()
	err := InsertOrUpdateOTP(tempToken, otp, otp_type)
	if err != nil {
		return err
	}
	return SendMail(tempToken, otp, otp_type)
}

func ValidateAndDeleteOTP(tempToken, otp, otp_type string) error {
	db := database.DB
	var deletedRows []model.OtpQueue
	db.Unscoped().Clauses(clause.Returning{}).Where("temp_token = ? AND otp = ? AND otp_type = ?", tempToken, otp, otp_type).Delete(&deletedRows)
	if len(deletedRows) == 0 {
		return fmt.Errorf("bad request")
	}
	return nil
}
