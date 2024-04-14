package checkversion

type EPGVersionsResponse struct {
	Version1 string `json:"version1"`
	Version2 string `json:"version2"`
	Error1   string `json:"error1,omitempty"`
	Error2   string `json:"error2,omitempty"`
}
