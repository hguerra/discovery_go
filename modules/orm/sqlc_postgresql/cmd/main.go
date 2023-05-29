package main

import (
	"context"
	"log"
	"os"
	"reflect"

	"github.com/hguerra/discovery_go/modules/orm/sqlc_postgresql/internal/infra/db/schema"
	// _ "github.com/lib/pq"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func run() error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := schema.New(conn)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, schema.CreateAuthorParams{
		Name: "Brian Kernighan",
		// Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
		Bio: pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))

	// list
	authorsByIDs, err := queries.ListAuthorsByIDs(ctx, []int32{1, 2})
	if err != nil {
		return err
	}
	log.Println(authorsByIDs)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
