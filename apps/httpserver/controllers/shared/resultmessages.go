package shared

import "encoding/json"

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewHTTPErrorJSON(code int, message string) string {
	httpError := HTTPError{Code: code, Message: message}
	errorJson, _ := json.Marshal(&httpError)
	return string(errorJson)
}

type HTTPOk struct {
	Message string `json:"message"`
}

type HTTPPostCreationOk struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}
