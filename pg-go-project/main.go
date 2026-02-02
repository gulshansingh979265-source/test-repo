package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	dsn := "postgres://testuser:testpass@127.0.0.1:5432/testdb?sslmode=disable"

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}
	defer conn.Close(context.Background())

	// Use time.Time instead of string
	var now time.Time
	err = conn.QueryRow(context.Background(), "SELE NOW()").Scan(&now)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Connected! DB time:", now)
	fmt.Println("✅ Formatted:", now.Format("2006-01-02 15:04:05"))
}
