package submitcard

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/milkaxq/bpcpayment/config"
	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
	"golang.org/x/net/html"
)

func SubmitCardToBank(apiClient string, submitCardRequest SubmitCardRequest) (SubmitCardResponse, error) {
	var (
		req *http.Request
		err error
	)

	if apiClient == "senagat" {
		req, err = createSenagatSubmitCardRequest(submitCardRequest)
	} else if apiClient == "halkbank" {
		req, err = createHalkBankSubmitRequest(submitCardRequest)
	}
	if err != nil {
		return SubmitCardResponse{}, errors.New(constants.ErrCreatingRequest + err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return SubmitCardResponse{}, errors.New(constants.ErrSengdingRequest + err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SubmitCardResponse{}, errors.New(constants.ErrReadingResponseBody + err.Error())
	}

	var submitCardResponse SubmitCardResponse

	err = json.Unmarshal(body, &submitCardResponse)
	if err != nil {
		return SubmitCardResponse{}, errors.New(constants.ErrReadingResponseBody + err.Error())
	}

	return submitCardResponse, nil
}

func createSenagatSubmitCardRequest(submitCardRequest SubmitCardRequest) (*http.Request, error) {
	urlParams := utils.StructToURLParams(submitCardRequest)
	fullUrl := fmt.Sprintf("https://%s%s", constants.Base.SenagatBaseURL, constants.SenagatConfirmPaymentURL) + "?" + urlParams
	req, err := http.NewRequest("POST", fullUrl, nil)
	if err != nil {
		return nil, err
	}
	return req, err
}

func createHalkBankSubmitRequest(submitCardRequest SubmitCardRequest) (*http.Request, error) {
	formData := url.Values{}
	formData.Add("MDORDER", submitCardRequest.MDORDER)
	formData.Add("$EXPIRY", submitCardRequest.EXPIRY)
	formData.Add("$PAN", submitCardRequest.PAN)
	formData.Add("TEXT", submitCardRequest.TEXT)
	formData.Add("$CVC", submitCardRequest.CVC)
	formData.Add("language", submitCardRequest.Language)
	encodedData := formData.Encode()
	fullUrl := fmt.Sprintf("https://%s%s", constants.Base.HalkBaseURL, constants.HalkBankConfirmPaymentURL)
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBufferString(encodedData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func getOTPPassword(submitCardResponse SubmitCardResponse, MDOrder, apiClient string) (string, error) {
	formData := url.Values{}
	formData.Add("MD", MDOrder)
	formData.Add("PaReq", submitCardResponse.PaReq)
	formData.Add("TermUrl", submitCardResponse.TermUrl)
	encodedData := formData.Encode()

	req, err := http.NewRequest("POST", submitCardResponse.AcsUrl, bytes.NewBufferString(encodedData))
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

	request_id := utils.FindRequestId(root)

	if apiClient == "senagat" {
		err = sendOTPForSenagat(request_id)
	} else if apiClient == "halkbank" {
		err = sendOTPForHalkBank(request_id)
	}
	if err != nil {
		return "", errors.New(constants.ErrSendingOTP + err.Error())
	}
	return request_id, nil
}

func sendOTPForHalkBank(requestID string) error {
	formData := url.Values{}
	formData.Add("request_id", requestID)
	formData.Add("authForm", "authForm")
	formData.Add("sendPasswordButton", "Send password")
	encodedData := formData.Encode()

	fullUrl := fmt.Sprintf("https://%s", constants.HalkBankOtpUrl)
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBufferString(encodedData))
	if err != nil {
		return errors.New(constants.ErrCreatingRequest + err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New(constants.ErrSengdingRequest + err.Error())
	}
	defer resp.Body.Close()

	return nil
}

func sendOTPForSenagat(requestID string) error {
	formData := url.Values{}
	formData.Add("request_id", requestID)
	formData.Add("sendButton", "Ugratmak")
	encodedData := formData.Encode()

	fullUrl := fmt.Sprintf("https://%s%s", constants.Base.SenagatBaseURL, constants.SenagatOTPURL)
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBufferString(encodedData))
	if err != nil {
		return errors.New(constants.ErrCreatingRequest + err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New(constants.ErrSengdingRequest + err.Error())
	}
	defer resp.Body.Close()

	return nil
}

func addOTPRequestID(orderID string, bankModel config.BankModel) error {
	data, _ := json.Marshal(bankModel)

	err := utils.CreateNewEntry(orderID, data)
	if err != nil {
		return err
	}
	return nil
}
