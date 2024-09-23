package handler

import (
	"strconv"
	"math/rand"
)

func GenerateOTP() string {
	otp:=""
	for i:=0;i<4;i++ {
		otp+=strconv.Itoa(rand.Intn(10))
	}
	return otp
}

// func TwoFA(tempToken string) error {
// 	otp:=GenerateOTP()

// }
