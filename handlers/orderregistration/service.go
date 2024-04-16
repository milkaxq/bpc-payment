package orderregistration

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/milkaxq/bpcpayment/constants"
)

func makeOrderRegistration(urlParams string) (OrderRegistrationResponse, error) {
	fullUrl := fmt.Sprintf("https://%s%s", constants.Base.BaseURL, constants.RegisterURL) + "?" + urlParams

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
