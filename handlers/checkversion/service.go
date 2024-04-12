package checkversion

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/milkaxq/bpcpayment/constants"
)

func checkEPGVersions(baseURL string) (EPGVersionsResponse, error) {
	url1 := fmt.Sprintf("https://%s/epg/version.do", baseURL)
	version1, err1 := checkEPGVersion(url1)
	if err1 != nil {
		return EPGVersionsResponse{}, errors.New(constants.ErrCheckingEPGVersion + "(URL 1) " + err1.Error())
	}

	url2 := fmt.Sprintf("https://%s/epg/version", baseURL)
	version2, err2 := checkEPGVersion(url2)
	if err2 != nil {
		return EPGVersionsResponse{}, errors.New(constants.ErrCheckingEPGVersion + "(URL 2) " + err2.Error())
	}

	return EPGVersionsResponse{
		Version1: version1,
		Version2: version2,
	}, nil
}

func checkEPGVersion(url string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", errors.New(constants.ErrCreatingRequest + err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(constants.ErrSengdingRequest + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(constants.ErrUnexpectedStatusCode + strconv.Itoa(resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(constants.ErrReadingResponseBdoy + err.Error())
	}

	return string(body), nil
}
