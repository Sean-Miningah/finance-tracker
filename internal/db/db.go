package db

import (
	"github.com/jmoiron/sqlx"
)

func NewPGStorage(datasource string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", datasource)
	if err != nil {
		return nil, err
	}
	return db, nil
}
