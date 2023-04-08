package app

import (
	"bot_svc/internal/config"
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func newStorage(conf *config.DbConf) *sqlx.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, conf.Driver, conf.Conn)
	if err != nil {
		log.Fatalf("unable open db: %v\n", err)
	}
	db.Exec(schema)
	return db
}

var schema = `
CREATE TABLE IF NOT EXISTS messages (
    uuid UUID PRIMARY,
    name TEXT NOT NULL,
    body TEXT NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);`
