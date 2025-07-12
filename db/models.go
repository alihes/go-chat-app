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
	Sender		string		`json:"sender"`
	ReceiverID	int
	Content		string		`json:"content"`
	Timestamp	time.Time	`json:"timestamp"`
}

func InsertMessage(ctx context.Context, senderID, receiverID int, content string) error {
	_,err := Pool.Exec(ctx, `
		INSERT INTO messages (sender_id, receiver_id, content)
		VALUES ($1, $2, $3)
		`, senderID, receiverID, content)
		fmt.Println(err)
		return err
}

func GetMessages(ctx context.Context, limit int) ([]Message, error) {
	row, err := Pool.Query(ctx,`
		SELECT messages.content, messages.timestamp, users.username
			FROM messages JOIN users ON messages.sender_id = users.id
				ORDER BY messages.timestamp DESC LIMIT $1
		`, limit)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var msgs []Message
	for row.Next() {
		var m Message
		var username string
		err := row.Scan(&m.Content, &m.Timestamp, &username)
		if err != nil{
			return nil, err
		}
		m.Sender = username
		msgs = append(msgs, m)
	}

	return msgs, nil
}