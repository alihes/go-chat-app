package db

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	ID				int
	Username		string
	PasswordHash	string
	Code			string
}

type Message struct {
	ID			int
	SenderID	string
	ReceiverID	int
	Content		string
	Timestamp	time.Time
}

func InsertMessage(ctx context.Context, senderID, receiverID int, content string) error {
	_,err := Pool.Exec(ctx, `
		INSERT INTO messages (sender_id, receiver_id, content)
		VALUES ($1, $2, $3)
		`, senderID, receiverID, content)
		fmt.Println(err)
		return err
}