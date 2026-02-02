package main

import (
"context"
"fmt"
"log"

"github.com/jackc/pgx/v5"
)

func main() {
dsn := "postgres://testuser:testpass@localhost:5432/testdb?sslmode=disable"

conn, err := pgx.Connect(context.Background(), dsn)
if err != nil {
log.Fatal("Connection failed:", err)
}
defer conn.Close(context.Background())

var now string
err = conn.QueryRow(context.Background(), "SELECT NOW()").Scan(&now)
if err != nil {
log.Fatal(err)
}

fmt.Println("Connected! DB time:", now)
}


