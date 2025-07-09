package chat

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alihes/go-chat-app/db"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {return true},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

func HandleConnections(w http.ResponseWriter, r *http.Request){
	ws, err := upgrader.Upgrade(w,r,nil)
	if err != nil{
		fmt.Println("websocket upgrade failed:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var msg string
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Read error:", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
		senderID, receiverID := 1,2
		go db.InsertMessage(context.Background(), senderID, receiverID, msg)

	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			client.WriteJSON(msg)
		}
	}
}