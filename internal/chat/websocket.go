package chat

import (
	"context"
	"fmt"
	"net/http"
	// "strings"
	"time"

	"github.com/alihes/go-chat-app/db"
	"github.com/gorilla/websocket"
)


type OutgoingMessages struct {
	Sender		string	`json:"sender"`
	Content		string	`json:"content"`
	Timestamp	string	`json:"timestamp"`
}


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {return true},
}

var clients = make(map[*websocket.Conn]bool)
var userSockets = make(map[string]*websocket.Conn)
var onlineUsers = make(chan []string)
var broadcast = make(chan OutgoingMessages)

func HandleConnections(w http.ResponseWriter, r *http.Request){


	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "missing user code", http.StatusUnauthorized)
		return
	}

	var user db.User
	err := db.Pool.QueryRow(context.Background(),
		`SELECT ID, username FROM users WHERE code = $1`,
		code).Scan(&user.ID, &user.Username)

	if err != nil {
		http.Error(w, "invalid user code", http.StatusUnauthorized)
		return
	}


	ws, err := upgrader.Upgrade(w,r,nil)
	if err != nil{
		fmt.Println("websocket upgrade failed:", err)
		return
	}
	userSockets[user.Username] = ws
	updateOnlineUserList()
	defer ws.Close()



	clients[ws] = true

	for {
		var msgText string
		err := ws.ReadJSON(&msgText)
		if err != nil {
			fmt.Println("Read error:", err)
			delete(clients, ws)
			delete(userSockets, user.Username)
			updateOnlineUserList()
			break
		}
		broadcast <- OutgoingMessages{
			Sender: user.Username,
			Content: msgText,
			Timestamp: time.Now().Format(time.RFC3339),
		}
		senderID, receiverID := user.ID,0

		go db.InsertMessage(context.Background(), senderID, receiverID, msgText)

	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                fmt.Println("Write error:", err)
                client.Close()
                delete(clients, client)
            }
		}
	}
}

func BroadcastOnlineUsers() {
	for{
		users := <-onlineUsers
		for _,ws := range userSockets {
			ws.WriteJSON(struct {
				Type	string	`json:"type"`
				Users	[]string	`json:"type"`
			}{
				Type:	"online_users",
				Users: users,
			})
		}
	}
}

func updateOnlineUserList() {
	users := make([]string, 0, len(userSockets))
	for username := range userSockets {
		users = append(users, username)
	}
	onlineUsers <- users
}