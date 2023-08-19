package room

import (
	"log"
	"net/http"
	"slices"
	"strconv"
)

func HandleRoomConnection(w http.ResponseWriter, r *http.Request) {
	roomId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		return
	}

	hub := newHub(roomId)
	go hub.run()

	upgrader.CheckOrigin = func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		allowedOrigins := []string{"http://localhost:63342"}
		return slices.Contains(allowedOrigins, origin)
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
