package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgresql://root:1234@localhost:5433/wattage_monitor?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("could not connect to database")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
