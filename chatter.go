package chatserver

import (
	"github.com/gorilla/websocket"
)

type Chatter struct {
	username string
	room *Room
	socket *websocket.Conn
	mailbox chan []byte //changed to []byte cuz that's what it seems like JSON will send?
}

func (c *Chatter) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *Chatter) write() {
	defer c.socket.Close()
	for out := range c.mailbox { //iteration will break when channel is closed
		err := c.socket.WriteMessage(websocket.TextMessage, out)
		if err != nil {
			return
		}
	}
}

// when someone goes to chat URL, they choose a room and username
// then a Chatter gets initialized with a new WS connection
// chat.go has a function to continuously listen to connection (connection.ReadMessage)
// when it receives a message from conn, put it in channel room.messages
// it could also receive stuff on the websocket connection??
// or just have the room post to a wall and the wall can be displayed in the browser??
// if its just posting to a wall and reading the wall, then what's the post of 2-way WS instead of HTTP?
// use WS "send()" method to send text to the server
// define a handler function on our WebSocket's "onmessage" property to do something with messages sent from the server

