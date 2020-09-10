package chatserver

type Message struct {
	sender string `json: "sender""`
	timestamp string `json: "timestamp"`
	message string `json: "message"`
}


