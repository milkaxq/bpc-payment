package orderregistration

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

func makeOrderRegistration(apiclient, urlParams string) (OrderRegistrationResponse, error) {
	var fullURL string
	if apiclient == "senagat" {
		fullURL = fmt.Sprintf("https://%s%s?", constants.Base.SenagatBaseURL, constants.SenagatRegisterURL) + urlParams
	} else if apiclient == "halkbank" {
		fullURL = fmt.Sprintf("https://%s%s?", constants.Base.HalkBaseURL, constants.HalkBankRegisterURL) + urlParams
	}

	req, err := http.NewRequest("POST", fullURL, nil)
	if err != nil {
		return OrderRegistrationResponse{}, errors.New(constants.ErrCreatingRequest + err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return OrderRegistrationResponse{}, errors.New(constants.ErrSengdingRequest + err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return OrderRegistrationResponse{}, errors.New(constants.ErrReadingResponseBody + err.Error())
	}

	var orderRegistrationResponse OrderRegistrationResponse

	err = json.Unmarshal(body, &orderRegistrationResponse)
	if err != nil {
		return OrderRegistrationResponse{}, errors.New(constants.ErrReadingResponseBody + err.Error())
	}

	if orderRegistrationResponse.ErrorMessage != "" {
		return OrderRegistrationResponse{}, errors.New(constants.ErrInvalidResposnseBody + orderRegistrationResponse.ErrorMessage)
	}

	return orderRegistrationResponse, nil
}

func createNewOrder(orderRegistrationRequest OrderRegistrationRequest, resp OrderRegistrationResponse) error {
	var bankModel config.BankModel = config.BankModel{
		ApiClient: orderRegistrationRequest.ApiClient,
		OrderID:   resp.OrderId,
	}
	data, err := json.Marshal(bankModel)
	if err != nil {
		return err
	}

	err = utils.CreateNewEntry(resp.OrderId, data)
	if err != nil {
		return err
	}

	return nil
}
