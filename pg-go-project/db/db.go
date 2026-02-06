package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"go-pg-test/student"
)

// EnsureStudentsTable creates the students table if it does not exist.
func EnsureStudentsTable(ctx context.Context, conn *pgx.Conn) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS students (
		id   SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age  INT  NOT NULL
	);`

	_, err := conn.Exec(ctx, createTableSQL)
	return err
}

// SeedStudents inserts some example student rows.
// It is safe to run multiple times; errors are logged but not returned.
func SeedStudents(ctx context.Context, conn *pgx.Conn) error {
	if _, err := conn.Exec(
		ctx,
		"INSERT INTO students (name, age) VALUES ($1, $2);",
		"Alice", 20,
	); err != nil {
		log.Printf("insert Alice failed (might be duplicate if rerun): %v", err)
	}

	if _, err := conn.Exec(
		ctx,
		"INSERT INTO students (name, age) VALUES ($1, $2);",
		"Bob", 22,
	); err != nil {
		log.Printf("insert Bob failed (might be duplicate if rerun): %v", err)
	}

	return nil
}

// GetAllStudents returns all students from the database.
func GetAllStudents(ctx context.Context, conn *pgx.Conn) ([]student.Student, error) {
	rows, err := conn.Query(ctx, "SELECT id, name, age FROM students ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studentsList []student.Student
	for rows.Next() {
		var s student.Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Age); err != nil {
			return nil, err
		}
		studentsList = append(studentsList, s)
	}

	return studentsList, nil
}
