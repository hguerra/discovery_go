package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func OpenDB(ctx context.Context, dsn string, setLimits bool) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if setLimits {
		log.Println("setting limits")
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)
		db.SetConnMaxLifetime(time.Minute * 30)
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err = db.PingContext(ctxWithTimeout)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func OpenDBX(ctx context.Context, dsn string, setLimits bool) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if setLimits {
		log.Println("setting limits")
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)
		db.SetConnMaxLifetime(time.Minute * 30)
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err = db.PingContext(ctxWithTimeout)
	if err != nil {
		return nil, err
	}

	return db, nil
}
