package model

import "encoding/json"

// Response Ответ от сервера chrome.
type Response struct {
	Id        int             `json:"id"`
	Result    json.RawMessage `json:"result,omitempty"`
	SessionID string          `json:"sessionId,omitempty"`
	Error     *ResponseError  `json:"error,omitempty"`
}

// ResponseError Ошибка сервера chrome.
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
