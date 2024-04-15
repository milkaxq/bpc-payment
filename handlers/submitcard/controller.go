package submitcard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
	"golang.org/x/net/html"
)

func SubmitCard(c *gin.Context) {
	var submitCardRequest SubmitCardRequest

	if err := c.BindJSON(&submitCardRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequestBody + err.Error()})
		return
	}

	urlParams := utils.StructToURLParams(submitCardRequest)

	fullUrl := fmt.Sprintf("https://%s/epg/rest/processform.do", constants.Base.BaseURL) + "?" + urlParams
	req, err := http.NewRequest("POST", fullUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var submitCardResponse SubmitCardResponse

	err = json.Unmarshal(body, &submitCardResponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	formData := url.Values{}
	formData.Add("MD", submitCardRequest.MDORDER)
	formData.Add("PaReq", submitCardResponse.PaReq)
	formData.Add("TermUrl", submitCardResponse.TermUrl)
	encodedData := formData.Encode()

	req, err = http.NewRequest("POST", submitCardResponse.AcsUrl, bytes.NewBufferString(encodedData))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	request_id := utils.FindRequestId(root)

	formData = url.Values{}
	formData.Add("request_id", request_id)
	formData.Add("sendButton", "Ugratmak")
	encodedData = formData.Encode()

	req, err = http.NewRequest("POST", "https://epg.senagatbank.com.tm/acs/api/3ds/form/otp", bytes.NewBufferString(encodedData))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	c.JSON(http.StatusOK, request_id)
}
