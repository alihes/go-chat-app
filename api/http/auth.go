package httpapi

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/alihes/go-chat-app/db"
	"golang.org/x/crypto/bcrypt"
	"honnef.co/go/tools/analysis/code"
)

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authResponse struct {
	Code string `json:"code"` //uuid
}


//POST /singup
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req authRequest
	
	//?
	var code string

	json.NewDecoder(r.Body).Decode(&req)

	hash,_ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	err := db.Pool.QueryRow(context.Background(),
		`INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING code`,
	req.Username, string(hash),).Scan(&code)

	if err != nil {
		http.Error(w, "signup failed", http.StatusBadRequest)
		return
	}

	json.NewEncoder(2).Encode(authResponse{Code: code})
}

//POST /login
func LoginHandler(w http.ResponseWriter,r *http.Request) {
	var req authRequest
	json.NewDecoder(r.Body).Decode(&req)

	var user db.User
	err := db.Pool.QueryRow(context.Background(),
		`SELECT id, password_hash, code FROM users WHERE username = $1`,
		req.Username).Scan(&user.ID, &user.PasswordHash, &user.Code)
	id err != nil{
		http.Error(w, "user not found", http.StatusBadRequest)
	}
}