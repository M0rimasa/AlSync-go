package myanimelist

import (
	"bytes"
	"fmt"
	"net/http"
)


var client = &http.Client{}

func convStatus(status string) int {
	switch status {
	case "CURRENT":
		return 1
	case "COMPLETED":
		return 2
	case "PAUSED":
		return 3
	case "DROPPED":
		return 4
	case "PLANNING":
		return 6
	default:
		return 1
	}
}

func sendJSON(url string, jsonData []byte, cookies string) (resCode int, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Cookie", cookies)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	resCode = resp.StatusCode
	return
}

type requestURL struct {
	mediaType string
	action string
}

func (r requestURL) String() string {
	return fmt.Sprintf("%v/%v/%v.json", baseURL, r.mediaType, r.action)
}
