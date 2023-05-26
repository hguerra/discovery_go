package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/hguerra/discovery_go/modules/orm/ent_postgresql/ent"
	"github.com/hguerra/discovery_go/modules/orm/ent_postgresql/ent/car"
	"github.com/hguerra/discovery_go/modules/orm/ent_postgresql/ent/user"

	// pg
	// _ "github.com/lib/pq"

	// pgx
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(28).
		SetName("Heitor").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// Create a new car with model "Tesla".
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	// Create a new car with model "Ford".
	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", ford)

	// Create a new user, and add it the 2 cars.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}

func QueryCars(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println("returned cars:", cars)

	// What about filtering specific cars.
	ford, err := a8m.QueryCars().
		Where(car.Model("Ford")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println(ford)
	return nil
}

func QueryCarUsers(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	// Query the inverse edge.
	for _, c := range cars {
		owner, err := c.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %w", c.Model, err)
		}
		log.Printf("car %q owner: %q\n", c.Model, owner.Name)
	}
	return nil
}

func main() {
	// Using pg
	// https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
	// url := fmt.Sprintf(
	// 	"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	// 	"localhost",
	// 	"5432",
	// 	"docker",
	// 	"docker",
	// 	"docker",
	// )
	// client, err := ent.Open("postgres", url)

	client := Open("postgresql://docker:docker@localhost:5432/docker?search_path=public&sslmode=disable")
	defer client.Close()

	ctx := context.Background()

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// CreateUser(ctx, client)
	// CreateCars(ctx, client)
	u, err := QueryUser(ctx, client)
	if err != nil {
		return
	}
	err = QueryCars(ctx, u)
	if err != nil {
		log.Fatal(err)
	}
	err = QueryCarUsers(ctx, u)
	if err != nil {
		log.Fatal(err)
	}
}
