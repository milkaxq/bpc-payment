package submitcard

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
)

func SubmitCard(c *gin.Context) {
	var submitCardRequest SubmitCardRequest

	if err := c.BindJSON(&submitCardRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequestBody + err.Error()})
		return
	}

	urlParams := utils.StructToURLParams(submitCardRequest)

	fmt.Println(urlParams)
	fullUrl := fmt.Sprintf("https://%s/epg/rest/processform.do", constants.Base.BaseURL) + "?" + urlParams
	req, err := http.NewRequest("POST", fullUrl, nil)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var submitCardResponse SubmitCardResponse

	err = json.Unmarshal(body, &submitCardResponse)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, submitCardResponse)
}
