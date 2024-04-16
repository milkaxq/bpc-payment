package constants

const (
	ErrInvalidRequestBody      = "invalid request body : "
	ErrCreatingRequest         = "error creating request: "
	ErrSengdingRequest         = "error sending request: "
	ErrUnexpectedStatusCode    = "unexpected status code: "
	ErrReadingResponseBody     = "error reading response body: "
	ErrCheckingEPGVersion      = "error checking EPG version: "
	ErrInvalidResposnseBody    = "invalid response body: "
	ErrParsingHtmlFromResponse = "couldn't parse html from response"
)

type HttpError struct {
	Error string `json:"error"`
}

// Constants for error codes
const (
	ErrorCodeNoError              = 0
	ErrorCodeProcessed            = 1
	ErrorCodeNotPaid              = 2
	ErrorCodeUnknownCurrency      = 3
	ErrorCodeMissingParameter     = 4
	ErrorCodeIncorrectValue       = 5
	ErrorCodeAccessDenied         = 6
	ErrorCodeSystemError          = 7
	ErrorCodeInvalidPaymentMethod = 14
)

// Map of error codes to corresponding error messages
var ErrorCodeToMessage = map[int]string{
	ErrorCodeNoError:              "",
	ErrorCodeProcessed:            "Order already processed or incorrect childId.",
	ErrorCodeNotPaid:              "Order registered but not paid.",
	ErrorCodeUnknownCurrency:      "Unknown currency.",
	ErrorCodeMissingParameter:     "Missing required parameter(s):",
	ErrorCodeIncorrectValue:       "Incorrect value for a request parameter or language code.",
	ErrorCodeAccessDenied:         "Access denied or other issues:",
	ErrorCodeSystemError:          "System error. Please try again later.",
	ErrorCodeInvalidPaymentMethod: "Invalid payment method provided.",
}
