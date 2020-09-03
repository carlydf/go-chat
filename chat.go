package chat

import (
	"github.com/gorilla/websocket"
)

type Chatter struct {
	username string
	room *Room
	conn *websocket.Conn
	outbox chan string //messages could be strings or could be structs with username, timestamp, string
}

// when someone goes to chat URL, they choose a room and username
// then a Chatter gets initialized with a new WS connection
// chat.go has a function to continuously listen to connection (connection.ReadMessage)
// when it receives a message from conn, put it in channel room.messages
// it could also receive stuff on the websocket connection??
// or just have the room post to a wall and the wall can be displayed in the browser??
// if its just posting to a wall and reading the wall, then what's the post of 2-way WS instead of HTTP?

