package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jaydeep-savaliya/otp/data"
)

func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var payload data.OTPData
		app.validateBody(c, &payload)
		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}
		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJson(c, err)
			return
		}
		app.writeJSON(c, http.StatusAccepted, "otp send successfully")
	}
}

func (app *Config) verifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var payload data.VerifyData
		app.validateBody(c, &payload)
		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}
		err := app.twilioverifyOTP(newData.User.PhoneNumber, newData.Code)
		if err != nil {
			app.errorJson(c, err)
			return
		}
		app.writeJSON(c, http.StatusAccepted, "otp verified successfully")
	}
}
