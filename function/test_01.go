package main

import "fmt"

type ApiError struct {
	Code    string
	Message string
}

func (e ApiError) String() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e ApiError) NewApiError(code, message string) *ApiError {
	return &ApiError{Code:code, Message:message}
}

func main() {
	var apiError = ApiError{Code:"001", Message:"흐흐"}
	fmt.Println(&apiError)
	fmt.Println(apiError)
}

