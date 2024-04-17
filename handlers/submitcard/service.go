package submitcard

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
	"golang.org/x/net/html"
)

func SubmitCardToBank(urlParams string) (SubmitCardResponse, error) {
	fullUrl := fmt.Sprintf("https://%s%s", constants.Base.SenagatBaseURL, constants.ConfirmPaymentURL) + "?" + urlParams
	req, err := http.NewRequest("POST", fullUrl, nil)
	if err != nil {
		return SubmitCardResponse{}, errors.New(constants.ErrCreatingRequest + err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

func GetOTPPassword(submitCardResponse SubmitCardResponse, MDOrder string) (string, error) {
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

	formData = url.Values{}
	formData.Add("request_id", request_id)
	formData.Add("sendButton", "Ugratmak")
	encodedData = formData.Encode()

	fullUrl := fmt.Sprintf("https://%s%s", constants.Base.SenagatBaseURL, constants.OTPURL)
	req, err = http.NewRequest("POST", fullUrl, bytes.NewBufferString(encodedData))
	if err != nil {
		return "", errors.New(constants.ErrCreatingRequest + err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return "", errors.New(constants.ErrSengdingRequest + err.Error())
	}
	defer resp.Body.Close()

	return request_id, nil
}
