package constants

const (
	ErrInvalidRequestBody   = "invalid request body : "
	ErrCreatingRequest      = "error creating request: "
	ErrSengdingRequest      = "error sending request: "
	ErrUnexpectedStatusCode = "unexpected status code: "
	ErrReadingResponseBdoy  = "error reading response body: "
	ErrCheckingEPGVersion   = "error checking EPG version: "
)

type HttpError struct {
	Error string `json:"error"`
}
