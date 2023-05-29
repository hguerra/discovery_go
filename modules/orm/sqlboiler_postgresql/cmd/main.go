package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/hguerra/discovery_go/modules/orm/sqlboiler_postgresql/internal/infra/db/schema"
	_ "github.com/lib/pq"

	// "github.com/jackc/pgx/v5"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(ctx context.Context) (*schema.User, error) {
	u := &schema.User{
		Name:     null.NewString("Heitor", true),
		LastName: null.NewString("Carneiro", true),
		Age:      null.NewInt(28, true),
	}

	err := u.InsertG(ctx, boil.Infer())
	if err != nil {
		return nil, err
	}
	return u, nil
}

func createBlogPost(ctx context.Context, author *schema.User) (*schema.BlogPost, error) {
	b := &schema.BlogPost{
		Title:    null.NewString("Title 1", true),
		Body:     null.NewString("Body 1", true),
		AuthorID: null.Int64From(author.ID),
	}

	err := b.InsertG(ctx, boil.Infer())
	if err != nil {
		return nil, err
	}
	return b, nil
}

func selectUsersWithPosts(ctx context.Context, authorID int64) {
	author, err := schema.Users(
		schema.UserWhere.ID.EQ(authorID),
		qm.Load(schema.UserRels.AuthorBlogPosts),
	).
		OneG(ctx)

	catch(err)
	fmt.Println(author)

	// posts, err := author.AuthorBlogPosts().AllG(ctx)
	// catch(err)
	// fmt.Println(posts)

	for _, p := range author.R.AuthorBlogPosts {
		fmt.Println(p)
	}
}

func main() {
	ctx := context.Background()
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	catch(err)
	defer db.Close()

	boil.SetDB(db)

	user, err := createUser(ctx)
	catch(err)
	fmt.Println(user)

	blog, err := createBlogPost(ctx, user)
	catch(err)
	fmt.Println(blog)

	selectUsersWithPosts(ctx, user.ID)
}
