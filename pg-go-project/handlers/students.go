package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"go-pg-test/db"
)

// StudentsHandler returns an http.HandlerFunc that lists all students from Postgres.
func StudentsHandler(conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		studentsList, err := db.GetAllStudents(r.Context(), conn)
		if err != nil {
			http.Error(w, "failed to query students", http.StatusInternalServerError)
			log.Printf("query students failed: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(studentsList); err != nil {
			log.Printf("encode students failed: %v", err)
		}
	}
}

