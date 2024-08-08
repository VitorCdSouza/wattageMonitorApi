package main

import (
	"context"
	"fmt"
	"os"

	db "github.com/VitorCdSouza/wattageMonitorApi/db/sqlc"
	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgresql://root:1234@localhost:5433/wattage_monitor?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	q := db.New(conn)

	user, err := q.GetUser(context.Background(), 1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetAuthor failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(user.UserEmail)
}
