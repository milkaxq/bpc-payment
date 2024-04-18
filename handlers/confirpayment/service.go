package confirpayment

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
	"golang.org/x/net/html"
)

func ConfirmPaymenRequest(confirPaymentRequest ConfirmPaymentRequest) (string, error) {
	formData := url.Values{}
	formData.Add("request_id", confirPaymentRequest.RequestID)
	formData.Add("passwordEdit", confirPaymentRequest.PasswordEdit)
	formData.Add("submitButton", "Tassyklamak")
	encodedData := formData.Encode()

	fullUrl := fmt.Sprintf("https://%s%s", constants.Base.SenagatBaseURL, constants.OTPURL)
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBufferString(encodedData))
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

	root, err := html.Parse(resp.Body)
	if err != nil {
		return "", errors.New(constants.ErrParsingHtmlFromResponse + err.Error())
	}
	paRes := utils.FindPaRes(root)

	if paRes != "" {
		return "", errors.New(constants.ErrInvalidRequestBody)
	}

	return paRes, nil
}

func FinishPayment(PaRes, MDOrder string) (string, error) {
	formData := url.Values{}
	formData.Add("MD", MDOrder)
	formData.Add("PaRes", PaRes)
	encodedData := formData.Encode()

	fullUrl := fmt.Sprintf("https://%s%s", constants.Base.SenagatBaseURL, constants.FinishURL)
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBufferString(encodedData))
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

	return "Succesfully Payed", nil
}
