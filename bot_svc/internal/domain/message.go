package domain

import (
	"bot_svc/pkg/proto"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	// SenderID int64
	ID       string `db:"id"`
	UserName string `db:"user_name"`
	Body     string `db:"body"`
	Created  time.Time
}

func (m *Message) Fill(msg *proto.Message) {
	m.ID = uuid.NewString()
	m.UserName = msg.GetUserName()
	m.Body = msg.GetBody()
	m.Created = time.Now()
}
