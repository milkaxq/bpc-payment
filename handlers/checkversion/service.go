package checkversion

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/milkaxq/bpcpayment/constants"
)

// @Description This function takes a base URL and constructs two separate URLs for EPG data version retrieval. It then calls checkEPGVersion on each URL and returns a combined response object.
func checkEPGVersions(baseURL string) EPGVersionsResponse {
	url1 := fmt.Sprintf("https://%s/epg/version.do", baseURL)
	version1, err1 := checkEPGVersion(url1)
	url2 := fmt.Sprintf("https://%s/epg/version", baseURL)
	version2, err2 := checkEPGVersion(url2)

	return EPGVersionsResponse{
		Version1: version1,
		Version2: version2,
		Error1:   err1.Error(),
		Error2:   err2.Error(),
	}
}

// @Description This function takes a URL and performs an HTTP GET request to retrieve the EPG version. It checks the status code and parses the response body to return the version string.
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
