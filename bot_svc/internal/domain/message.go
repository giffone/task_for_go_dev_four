package domain

import (
	"bot_svc/pkg/proto"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	// SenderID int64
	ID       string `db:"id"`
	UserName string `db:"user_name"`
	Body     string `db:"body"`
	Created  time.Time
}

func (n *Note) Fill(msg *proto.NewNote) {
	n.ID = uuid.NewString()
	n.UserName = msg.GetUserName()
	n.Body = msg.GetBody()
	n.Created = time.Now()
}

func (n *Note) Validate() error {
	if n.UserName == "" {
		return ErrUserName
	}
	if n.Body == "" {
		return ErrTextBody
	}
	return nil
}
