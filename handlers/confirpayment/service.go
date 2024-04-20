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

func ConfirmPaymenRequest(confirPaymentRequest ConfirmPaymentRequest, apiClient string) (string, error) {
	var (
		resp *http.Response
		err  error
	)
	if apiClient == "senagat" {
		resp, err = confirmSenagatRequest(confirPaymentRequest)
	} else if apiClient == "halkbank" {
		resp, err = confirmHalkBankRequest(confirPaymentRequest)
	}
	if err != nil {
		return "", errors.New(constants.ErrConfirmPayment + err.Error())
	}

	defer resp.Body.Close()
	root, err := html.Parse(resp.Body)
	if err != nil {
		return "", errors.New(constants.ErrParsingHtmlFromResponse + err.Error())
	}
	paRes := utils.FindPaRes(root)

	if paRes == "" {
		return "", errors.New(constants.ErrInvalidRequestBody + "Invalid PaRes")
	}

	return paRes, nil
}

func confirmSenagatRequest(confirPaymentRequest ConfirmPaymentRequest) (*http.Response, error) {
	formData := url.Values{}
	formData.Add("request_id", confirPaymentRequest.RequestID)
	formData.Add("passwordEdit", confirPaymentRequest.PasswordEdit)
	formData.Add("submitButton", "Tassyklamak")
	encodedData := formData.Encode()

	fullUrl := fmt.Sprintf("https://%s%s", constants.Base.SenagatBaseURL, constants.SenagatOTPURL)
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBufferString(encodedData))
	if err != nil {
		return nil, errors.New(constants.ErrCreatingRequest + err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(constants.ErrSengdingRequest + err.Error())
	}

	return resp, nil
}

func confirmHalkBankRequest(confirPaymentRequest ConfirmPaymentRequest) (*http.Response, error) {
	formData := url.Values{}
	formData.Add("authForm", "authForm")
	formData.Add("request_id", confirPaymentRequest.RequestID)
	formData.Add("pwdInputVisible", confirPaymentRequest.PasswordEdit)
	formData.Add("submitPasswordButton", "Submit")
	encodedData := formData.Encode()

	fullUrl := fmt.Sprintf("https://%s", constants.HalkBankOtpUrl)
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBufferString(encodedData))
	if err != nil {
		return nil, errors.New(constants.ErrCreatingRequest + err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(constants.ErrSengdingRequest + err.Error())
	}

	return resp, nil
}

func finishPayment(PaRes, MDOrder, apiClient string) (string, error) {
	formData := url.Values{}
	formData.Add("MD", MDOrder)
	formData.Add("PaRes", PaRes)
	encodedData := formData.Encode()

	var fullURL string
	if apiClient == "senagat" {
		fullURL = fmt.Sprintf("https://%s%s", constants.Base.SenagatBaseURL, constants.SenagatFinishURL)
	} else if apiClient == "halkbank" {
		fullURL = fmt.Sprintf("https://%s%s", constants.Base.HalkBaseURL, constants.HalkBankFinishURL)
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

	return "Succesfully Payed", nil
}
