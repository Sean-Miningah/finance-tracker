package main

import (
	"log"

	"finance-crud-app/cmd/api"
	"finance-crud-app/internal/db"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Server struct {
	db  *sqlx.DB
	mux *mux.Router
}

func NewServer(db *sqlx.DB, mux *mux.Router) *Server {
	return &Server{
		db:  db,
		mux: mux,
	}
}

func main() {
	connStr := "postgres://postgres:Password123@localhost:5432/crud_db?sslmode=disable"

	db, err := db.NewPGStorage(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := api.NewAPIServer(":8085", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
