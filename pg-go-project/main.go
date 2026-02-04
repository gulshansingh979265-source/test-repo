package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {

	// ✅ PostgreSQL DSN
	dsn := "postgres://testuser:testpass@127.0.0.1:5432/testdb?sslmode=disable"

	// ✅ Connect Database
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal("❌ Connection failed:", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("✅ Connected Successfully!")

	// ✅ Check DB Time
	var now time.Time
	err = conn.QueryRow(context.Background(), "SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal("❌ Time Query Error:", err)
	}

	fmt.Println("DB Time:", now.Format("2006-01-02 15:04:05"))

	// ====================================================
	// ✅ STEP 1: Create Student Table
	// ====================================================
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INT NOT NULL
	);`

	_, err = conn.Exec(context.Background(), createTableSQL)
	if err != nil {
		log.Fatal("❌ Table Create Error:", err)
	}

	fmt.Println("✅ Student Table Created!")

	// ====================================================
	// ✅ STEP 2: Insert Student Data
	// ====================================================
	insertSQL := `
	INSERT INTO students (name, age)
	VALUES ($1, $2);`

	_, err = conn.Exec(context.Background(), insertSQL, "Amit", 20)
	if err != nil {
		log.Fatal("❌ Insert Error:", err)
	}

	fmt.Println("✅ Student Inserted Successfully!")

	// ====================================================
	// ✅ STEP 3: Fetch Students
	// ====================================================
	fmt.Println("\n📌 Student List:")

	rows, err := conn.Query(context.Background(), "SELECT id, name, age FROM students")
	if err != nil {
		log.Fatal("❌ Select Error:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int

		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %d | Name: %s | Age: %d\n", id, name, age)
	}

	fmt.Println("\n✅ Done!")
}

