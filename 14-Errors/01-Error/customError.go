package main

import (
	"fmt"
	"io"
	"net/http"
)

type HttpError struct {
	Message    string
	StatusCode int
}

func main() {
	resp, err := GetRequest("")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
func (error HttpError) Error() string {
	return fmt.Sprintf("Http error occurred: %d  %s", error.StatusCode, error.Message)
}
func NewHttpError(status int, message string) *HttpError {
	return &HttpError{Message: message, StatusCode: status}
}
func GetRequest(url string) (string, error) {
	if len(url) == 0 {
		return "", NewHttpError(400, "url is empty")
	}
	response, err := http.Get(url)
	if err != nil {
		return "", NewHttpError(500, "error occurred")
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", NewHttpError(200, "body read error")
	}
	return string(responseBody), nil
}

