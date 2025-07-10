package main

import (
	"fmt"
	"net/http"

	"github.com/alihes/go-chat-app/config"
	"github.com/alihes/go-chat-app/internal/chat"
	"github.com/alihes/go-chat-app/db"
	
)
import httpapi "github.com/alihes/go-chat-app/api/http"

func main() {
	cfg := config.Load()

	err := db.Connect()
	if err != nil{
		panic(err)
	}

	go chat.HandleMessages()

	fs := http.FileServer(http.Dir("web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/",fs))


	http.HandleFunc("/", index)

	http.HandleFunc("/ws", chat.HandleConnections)

	http.HandleFunc("/signup", httpapi.SignupHandler)
	http.HandleFunc("/login", httpapi.LoginHandler)

	fmt.Println("Server listening on https://127.0.0.1:433")
	//err := http.ListenAndServe(":8080", nil)
	err = http.ListenAndServeTLS(cfg.Port, cfg.CertFile, cfg.KeyFile, nil)
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "secure chat server is running on HTTPS!")
	http.ServeFile(w, r, "web/index.html")
}
