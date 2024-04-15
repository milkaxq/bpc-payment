package submitcard

type SubmitCardRequest struct {
	MDORDER  string `json:"MDORDER"`
	EXPIRY   string `json:"$EXPIRY"`
	PAN      string `json:"$PAN"`
	TEXT     string `json:"TEXT"`
	CVC      string `json:"$CVC"`
	Language string `json:"language"`
	MM       string `json:"MM"`
	YYYY     string `json:"YYYY"`
}

type SubmitCardResponse struct {
	Info      string `json:"info"`
	AcsUrl    string `json:"acsUrl"`
	PaReq     string `json:"paReq"`
	TermUrl   string `json:"termUrl"`
	ErrorCode int    `json:"errorCode"`
}
