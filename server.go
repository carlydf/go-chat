package chat

import (
	"net/http"
)

// has a function that runs constantly in the background
// starts HTTP server with all the common rooms and halls in the house as Room structs
// use http.FileServer to serve static HTML files?

// for name in (list of room names)
//	r = new room with that name
//	http.Handle("/chat/"+name, r) ideally if i was hosting this on our web server it would be "web/chat/room-name" as the URL
//	r.Run()

// HTTP connection would be upgraded to a WS connection in room.go ?? maybe ?? or in the for loop here??
