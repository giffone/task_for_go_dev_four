package storage

import (
	"bot_svc/internal/domain"
	"bot_svc/internal/service"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

var (
	inserQuery = `INSERT INTO messages (id, user_name, body, created)
	VALUES (:id, :user_name, :body, :created)`
)

func New(db *sqlx.DB) service.Storage {
	return &storage{db: db}
}

type storage struct {
	db *sqlx.DB
}

func (s *storage) Insert(ctx context.Context, msg *domain.Note) (string, error) {
	if _, err := s.db.NamedExecContext(ctx, inserQuery, msg); err != nil {
		log.Printf("db: insert: exec: %v", err)
		return "", err
	}
	return msg.ID, nil
}

func (s *storage) Update(ctx context.Context, msg *domain.Note) bool {
	return false
}

func (s *storage) Delete(ctx context.Context, msg *domain.Note) bool {
	return false
}
