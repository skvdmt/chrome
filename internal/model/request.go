package model

import (
	"encoding/json"

	"github.com/skvdmt/chrome/internal/devtools/types/target"
)

// RequestOption Опция запроса.
type RequestOption func(r *Request) error

// Request Запрос к серверу chrome.
type Request struct {
	Id        int               `json:"id"`
	SessionId *target.SessionId `json:"sessionId,omitempty"`
	Method    string            `json:"method"`
	Params    json.RawMessage   `json:"params,omitempty"`
}

// Json Формат.
func (r *Request) Json() []byte {
	j, _ := json.Marshal(r)
	return j
}

// WithSessionId С id сессиии.
func WithSessionId(sid *target.SessionId) RequestOption {
	return func(r *Request) error {
		r.SessionId = sid
		return nil
	}
}
