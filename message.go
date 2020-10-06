package chatserver

import (
	"encoding/json"
)

type Message struct {
	sender string `json: "sender""`
	timestamp string `json: "timestamp"`
	message string `json: "message"`
}

// FromJSON created a new Message struct from given JSON
func FromJSON(jsonInput []byte) (message *Message) {
	json.Unmarshal(jsonInput, &message)
	return
}


