package chat

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

type Room struct {
	name string
	chatters map[*Chatter]bool //defines who is "online" could be used to forward messages to WS conns of chatters who are in the room
	join chan *Chatter //listens for chatters who want to join, adds them to online chatters
	leave chan *Chatter //listens for chatters who want to leave, removes them from online chatters
	forward chan []byte //holds incoming msgs that should be forwarded to other clients
	logfile string //path to log file? that could then post the rooms chat history as the intro screen of web/chat/room-name
}

func NewRoom(name string) *Room {
	r := Room{name: name}
	r.chatters = make(map[*Chatter]bool)
	r.join = make(chan *Chatter)
	r.leave = make(chan *Chatter)
	r.forward = make(chan []byte)
	r.logfile = "chatlogs/" + name //not final location, todo
	return &r
}

func (r *Room) Run() {
	file, err := os.OpenFile(r.logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file) //does log know about this output file in all functions in this file now?
	log.Printf("running chat room %s", r.name)
	for {
		select {
		case joiner := <-r.join:
			r.chatters[joiner] = true
		case leaver := <-r.leave:
			r.chatters[leaver] = false
		case msg := <-r.forward:
			//send msg to mailbox of all chatters in the room
			for chatter := range r.chatters {
				chatter.mailbox <- msg
			}
			//log msg
			log.Println("info to print from msg, need to convert it out of []byte?")
		}
	}
}

const messageBufferSize = 1024

var upgrader = &websocket.Upgrader{
	ReadBufferSize: messageBufferSize,
	WriteBufferSize: messageBufferSize}

//this is an http handler, will be HandleFunc-ed in server.go
func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// determine whether an incoming req from a diff domain is allowed to connect
	// what happens when this isn't here?
	// why do we do this b4 upgrading connection?
	upgrader.CheckOrigin = func(req *http.Request) bool { return true }
	// figure out how to get the username from the HTTP Request and how to prompt the client to enter a username
	name := req.FormValue("Username") //from a stackexchange that used a JSON form it seems
	// Upgrade takes pointer to HTTP Request and returns pointer to a WS connection, or an error
	// each new HTTP Request corresponds to a new chatter, so create one here
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
	}

	newChatter := &Chatter{
		username: name,
		room:     r,
		socket:   ws,
		mailbox:  make(chan []byte),
	}
	r.join <- newChatter
	defer r.chatterLeave(newChatter)
	go newChatter.write() //goroutine so read() can run concurrently
	newChatter.read() //direct instead of goroutine, will block operations
}

func (r *Room) chatterLeave(chatter *Chatter) {
	r.leave <- chatter
}
