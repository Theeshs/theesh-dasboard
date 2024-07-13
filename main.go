package main

import (
	"context"
	"fmt"
	"log"
	"theedashboard/ent"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open(dialect.Postgres, "postgres://postgres:mysecretpassword@127.0.0.1:5432/thee.me?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	example(ctx, client)
}

func example(ctx context.Context, client *ent.Client) {
	// list all users
	users, err := client.User.Query().All(ctx)
	fmt.Println("Users:", users)

	if err != nil {
		log.Fatalf("Unable to fetch: %v", err)
	}

}
