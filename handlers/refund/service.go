package refund

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/milkaxq/bpcpayment/config"
	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
)

func makeRefundRequest(refundsRequest RefundRequestBody, bankModel config.BankModel) (string, error) {
	var fullURL string

	urlParams := utils.StructToURLParams(refundsRequest)

	if bankModel.ApiClient == "senagat" {
		fullURL = fmt.Sprintf("https://%s%s?", constants.Base.SenagatBaseURL, constants.SenagatRefundURL) + urlParams
	} else if bankModel.ApiClient == "halkbank" {
		fullURL = fmt.Sprintf("https://%s%s?", constants.Base.HalkBaseURL, constants.HalkBankRefundURL) + urlParams
	}

	resp, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var bankResponse RefundBankResponse
	err = json.Unmarshal(bodyBytes, &bankResponse)
	if err != nil {
		return "", err
	}

	if bankResponse.ErrorCode != "0" {
		return "", errors.New(bankResponse.ErrorMessage)
	}

	return "Successfully refunded", err
}
