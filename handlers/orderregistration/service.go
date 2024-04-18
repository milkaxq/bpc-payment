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

func makeOrderRegistration(urlParams string) (OrderRegistrationResponse, error) {
	fullUrl := fmt.Sprintf("https://%s%s?", constants.Base.SenagatBaseURL, constants.RegisterURL) + urlParams

	req, err := http.NewRequest("POST", fullUrl, nil)
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

	if orderRegistrationResponse.ErrorCode != 0 {
		return OrderRegistrationResponse{}, errors.New(constants.ErrInvalidResposnseBody + orderRegistrationResponse.getErrorString())
	}

	return orderRegistrationResponse, nil
}

func createNewOrder(orderRegistrationRequest OrderRegistrationRequest, resp OrderRegistrationResponse) error {
	var bankModel config.BankModel = config.BankModel{
		Username: orderRegistrationRequest.Username,
		Password: orderRegistrationRequest.Password,
		OrderID:  resp.OrderId,
	}
	data, err := json.Marshal(bankModel)
	if err != nil {
		return err
	}

	err = utils.CreateNewEntry(resp.OrderId, data, 60)
	if err != nil {
		return err
	}

	return nil
}
