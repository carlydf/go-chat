package chat

import (
	"net/http"
)

var roomNames = []string{
	"10s",
	"20s",
	"30s",
	"40s",
	"100s",
	"200s",
	"300s",
	"study-room",
	"itos",
	"red-room",
	"common-room",
	"balc",
	"kitchen",
}

func StartServer() {
	//make an HTTP handler for each room
	for _, roomName := range roomNames {
		r := NewRoom(roomName)
		http.HandleFunc("/chat/" + roomName, r.ServeHTTP) //sets up the default router in the net/http pkg
	}
	http.ListenAndServe(":8090", nil) //nil says use default router we just set up
}