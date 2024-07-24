package testutils

import (
	"finance-crud-app/internal/db"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func init() {
	// migrate := flag.Bool("migrate", false, "perform database migration")
	// flag.Parse()

	var err error
	connStr := "postgres://postgres:Password123@localhost:5432/crud_db?sslmode=disable"

	DB, err = db.NewPGStorage(connStr)
	if err != nil {
		log.Fatalf("Error connecting to TestingDB: %v", err)
	}

	// Perform MIgration on TestDb
	// if *migrate {
	// 	log.Printf("Migrate run %v", migrate)
	// }
}
