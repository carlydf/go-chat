package chatserver

import (
	"log"
	"net/http"
	"path/filepath"
	"html/template"
	//"github.com/gorilla/websocket"
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

//make a struct to help handle html templates
//this struct satisfies the http.Handler interface
//thus it can be passed to http.Handle(pattern string, myCustomHandler)
//http.HandleFunc(pattern string, <Handler func such as ServeHTTP>) would also work, but whatever
type templateHandler struct {
	filename string
	templ *template.Template
}
//give it a method that is a handler func
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r  *http.Request) {
	templatePath := filepath.Join("/Users/carlydefrondeville/go/src/chatserver/static", t.filename)
	t.templ = template.Must(template.ParseFiles(templatePath))
	t.templ.Execute(w, nil)
}

func StartServer() {
	//make a template struct for the start page html template
	http.Handle("/", &templateHandler{filename: "index.html"})
	//make an HTTP handler for each room
	for _, roomName := range roomNames {
		r := NewRoom(roomName)
		http.HandleFunc("/chat/" + roomName, r.ServeHTTP) //sets up the default router in the net/http pkg
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}