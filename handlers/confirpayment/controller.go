package confirpayment

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
	"golang.org/x/net/html"
)

func ConfirmPayment(c *gin.Context) {
	var confirPaymentRequest ConfirmPaymentRequest

	if err := c.BindJSON(&confirPaymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequestBody + err.Error()})
		return
	}

	formData := url.Values{}
	formData.Add("request_id", confirPaymentRequest.RequestID)
	formData.Add("passwordEdit", confirPaymentRequest.PasswordEdit)
	formData.Add("submitButton", "Tassyklamak")
	encodedData := formData.Encode()

	fmt.Println(confirPaymentRequest)
	req, err := http.NewRequest("POST", "https://epg.senagatbank.com.tm/acs/api/3ds/form/otp", bytes.NewBufferString(encodedData))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	paRes := utils.FindPaRes(root)
	formData = url.Values{}
	formData.Add("MD", confirPaymentRequest.MDORDER)
	formData.Add("PaRes", paRes)
	encodedData = formData.Encode()

	fmt.Println(paRes)
	req, err = http.NewRequest("POST", "https://epg.senagatbank.com.tm/payments/rest/finish3ds.do", bytes.NewBufferString(encodedData))
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

	c.JSON(http.StatusOK, "Succesfully Payed")
}
