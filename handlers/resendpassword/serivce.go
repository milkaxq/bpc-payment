package resendpassword

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/milkaxq/bpcpayment/constants"
)

func resendOTPRequest(requestID, apiClient string) (string, error) {
	formData := url.Values{}
	formData.Add("request_id", requestID)
	formData.Add("resendButton", "Kody gaýtadan ugratmak")
	formData.Add("passwordEdit", "")
	encodedData := formData.Encode()

	var fullURL string
	if apiClient == "senagat" {
		fullURL = fmt.Sprintf("https://%s%s", constants.Base.SenagatBaseURL, constants.SenagatOTPURL)
	} else if apiClient == "halkbank" {
		fullURL = fmt.Sprintf("https://%s%s", constants.Base.HalkBaseURL, constants.HalkBankOtpUrl)
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
