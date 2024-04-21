package resendpassword

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/milkaxq/bpcpayment/config"
	"github.com/milkaxq/bpcpayment/constants"
)

func resendOTPRequest(bankModel config.BankModel) (string, error) {
	formData := url.Values{}
	if bankModel.ApiClient == "senagat" {
		formData.Add("request_id", bankModel.OTPRequestID)
		formData.Add("resendButton", "Kody gaýtadan ugratmak")
		formData.Add("passwordEdit", "")
	} else if bankModel.ApiClient == "halkbank" {
		formData.Add("authForm", "authForm")
		formData.Add("request_id", bankModel.OTPRequestID)
		formData.Add("pwdInputVisible", "")
		formData.Add("resendPasswordLink", "resendPasswordLink")
	}

	encodedData := formData.Encode()

	var fullURL string
	if bankModel.ApiClient == "senagat" {
		fullURL = fmt.Sprintf("https://%s%s", constants.Base.SenagatBaseURL, constants.SenagatOTPURL)
	} else if bankModel.ApiClient == "halkbank" {
		fullURL = fmt.Sprintf("https://%s", constants.HalkBankOtpUrl)
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewBufferString(encodedData))
	if err != nil {
		return "", errors.New(constants.ErrCreatingRequest + err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(constants.ErrSengdingRequest + err.Error())
	}
	defer resp.Body.Close()

	return "Succesfully Resended OTP", nil
}
