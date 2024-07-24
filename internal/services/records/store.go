package records

import (
	"finance-crud-app/internal/types"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type RecordsStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *RecordsStore {
	return &RecordsStore{db: db}
}

func (s *RecordsStore) GetRecordById(id string) (types.Record, error) {
	recordId, err := strconv.Atoi(id)
	if err != nil {
		return types.Record{}, fmt.Errorf("error converting id to int: %v", err)
	}

	record := types.Record{}
	err = s.db.Select(&record, "SELECT * FROM records where id = ?", recordId)
	if err != nil {
		return types.Record{}, fmt.Errorf("error retrieving record: %v", err)
	}

	return record, nil
}

func (s *RecordsStore) GetUserRecords(userId string) ([]types.Record, error) {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return nil, fmt.Errorf("error converting id to int: %v", err)
	}

	records := []types.Record{}

	err = s.db.Select(&records, "SELECT * FROM records WHERE userId = ?", userIdInt)
	if err != nil {
		return nil, fmt.Errorf("error retrieving records: %v", err)
	}

	return records, nil
}

func (s *RecordsStore) CreateUserRecord(userId string, record types.Record) error {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return fmt.Errorf("error converting id to int: %v", err)
	}

	_, err = s.db.Exec("INSERT INTO records (description, category, amount, userId) VALUES (?, ?, ?, ?)",
		record.Description, record.Category, record.Amount, userIdInt)
	if err != nil {
		return err
	}

	return nil
}

func (s *RecordsStore) GetUserRecordByCategory(userId string, category string) ([]types.Record, error) {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return nil, fmt.Errorf("error converting id to int: %v", err)
	}

	records := []types.Record{}

	query := `SELECT * FROM records WHERE userId = ? AND category = ? ORDER BY createdAt`
	err = s.db.Select(&records, query, userIdInt, category)
	if err != nil {
		return nil, fmt.Errorf("error retrieving records: %v", err)
	}

	return records, nil
}

func (s *RecordsStore) UpdateRecord(recordId string, update types.Record) error {
	return nil
}

func (s *RecordsStore) DeleteRecord(recordId string) error {
	return nil
}
