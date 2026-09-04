package model

import (
	"encoding/json"

	"github.com/skvdmt/chrome/internal/devtools/types/target"
)

// Event Событие.
type Event struct {
	Method    string           `json:"method"`
	Params    json.RawMessage  `json:"params"`
	SessionId target.SessionId `json:"sessionId"`
}
