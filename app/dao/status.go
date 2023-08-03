package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	status struct {
		db *sqlx.DB
	}
)

func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

func (s *status) CreateStatus(status *object.Status) error {
	_, err := s.db.Exec("INSERT INTO status (id, account_id, content) VALUES (?, ?, ?)",
		status.ID, status.AccountID, status.Content)
	if err != nil {
		return fmt.Errorf("failed to insert status into db: %w", err)
	}

	return nil
}

func (s *status) GetStatus(statusID int64) (*object.Status, error) {
	entity := new(object.Status)
	err := s.db.QueryRowx("SELECT * FROM status WHERE id = ?", statusID).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil
}
