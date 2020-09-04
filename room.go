package chat

import (
	"log"
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

func (r *Room) Run() {
	file, err := os.OpenFile(r.logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
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

//should have a main function to "run" the room
//use select statement to listen to join, leave, and messages channels and take actions as stuff comes in

