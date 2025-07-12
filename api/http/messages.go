package httpapi

import (
	"encoding/json"
	"net/http"
	"context"

	"github.com/alihes/go-chat-app/db"
)


func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	msgs, err := db.GetMessages(context.Background(), 50)
	if err != nil{
		http.Error(w, "could not fetch messages", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(msgs)
}