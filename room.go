package chat

import (
	"net/http"
	"github.com/gorilla/websocket"
	"net/http"
)

Room struct {
	name string
	chatters map[*Chatter]bool //defines who is "online" could be used to forward messages to WS conns of chatters who are in the room
	chatlog string //path to log file? that posts the chat history to a URL?
	join chan *Chatter //listens for chatters who want to join, adds them to online chatters
	leave chan *Chatter //listens for chatters who want to leave, removes them from online chatters
	messages chan string //listens for message strings (or message objects?) to forward to log and maybe to online WS connections
}

//should have a main function to "run" the room
//use select statement to listen to join, leave, and messages channels and take actions as stuff comes in

