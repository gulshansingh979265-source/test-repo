package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go-pg-test/db"
	"go-pg-test/handlers"

	"github.com/jackc/pgx/v5"
)

func main() {
	// PostgreSQL DSN (adjust if your local config is different)
	dsn := "postgres://testuser:testpass@localhost:5432/testdb?sslmode=disable"

	ctx := context.Background()

	// Connect to PostgreSQL
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer conn.Close(ctx)

	// Ensure students table exists
	if err := db.EnsureStudentsTable(ctx, conn); err != nil {
		log.Fatalf("create table failed: %v", err)
	}

	// Insert example students into PostgreSQL
	if err := db.SeedStudents(ctx, conn); err != nil {
		log.Printf("seeding students failed: %v", err)
	}

	// /students GET handler to list all students from PostgreSQL (defined in handlers package)
	http.HandleFunc("/students", handlers.StudentsHandler(conn))

	// Simple HTTP server on port 3000
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Simple Go server running on :3000")
	})

	// /hello handler in separate package
	http.HandleFunc("/hello", handlers.HelloHandler)

	log.Println("Server listening on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
