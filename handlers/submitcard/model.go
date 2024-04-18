package submitcard

type SubmitCardRequest struct {
	// Another meaning of order_id
	MDORDER string `json:"MDORDER"`
	// should be in form like YYYYMM example 202709
	EXPIRY string `json:"$EXPIRY"`
	// PAN is a number on card 12 digit code
	PAN string `json:"$PAN"`
	// Text is card holder name and surname
	TEXT string `json:"TEXT"`
	// CVC is secure code on back side of your card
	CVC string `json:"$CVC"`
	// language is language
	Language string `json:"language"`
}

type SubmitCardResponse struct {
	Info      string `json:"info"`
	AcsUrl    string `json:"acsUrl"`
	PaReq     string `json:"paReq"`
	TermUrl   string `json:"termUrl"`
	ErrorCode int    `json:"errorCode"`
}
