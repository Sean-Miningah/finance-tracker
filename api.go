package main

import (
	"database/sql"
	"encoding/json"
	"finance-crud-app/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *App) getRecords(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r)
	rows, err := a.db.Query("SELECT id, name, value FROM records")
	if err != nil {
		http.Error(w, "Failed to query records", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var records []model.Record
	for rows.Next() {
		var record model.Record
		if err := rows.Scan(&record.ID, &record.Name, &record.Value); err != nil {
			http.Error(w, "Failed to scan record", http.StatusInternalServerError)
			return
		}
		records = append(records, record)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating over rows", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(records); err != nil {
		http.Error(w, "Failed to encode records", http.StatusInternalServerError)
	}
}

func (a *App) getRecordByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid record ID", http.StatusBadRequest)
		return
	}

	var record model.Record
	err = a.db.QueryRow("SELECT id, name, value FROM records WHERE id = $1", id).Scan(&record.ID, &record.Name, &record.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Record not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to query record", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(record); err != nil {
		http.Error(w, "Failed to encode record", http.StatusInternalServerError)
	}
}
