package checkstatus

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
)

func CheckOrderStatus(orderStatusRequest OrderStatusRequest, apiClient string) (OrderStatusResponse, error) {
	urlParams := utils.StructToURLParams(orderStatusRequest)
	var fullURL string
	if apiClient == "senagat" {
		fullURL = fmt.Sprintf("https://%s%s?", constants.Base.SenagatBaseURL, constants.SenagatOrderStatusURL) + urlParams
	} else if apiClient == "halkbank" {
		fullURL = fmt.Sprintf("https://%s%s?", constants.Base.HalkBaseURL, constants.HalkBankOrderStatusURL) + urlParams
	}

	req, err := http.NewRequest("POST", fullURL, nil)
	if err != nil {
		return OrderStatusResponse{}, errors.New(constants.ErrCreatingRequest + err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return OrderStatusResponse{}, errors.New(constants.ErrSengdingRequest + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return OrderStatusResponse{}, errors.New(constants.ErrReadingResponseBody + err.Error())
	}

	var orderStatusResponse OrderStatusResponse

	err = json.Unmarshal(body, &orderStatusResponse)
	if err != nil {
		return OrderStatusResponse{}, errors.New(constants.ErrReadingResponseBody + err.Error())
	}
	errorCode, _ := strconv.Atoi(orderStatusResponse.ErrorCode)
	if errorCode != 0 {
		return OrderStatusResponse{}, errors.New(constants.ErrInvalidResposnseBody + orderStatusResponse.getErrorString())
	}

	return orderStatusResponse, nil
}
