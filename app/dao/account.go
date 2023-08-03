package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}

func (r *account) CreateAccount(account *object.Account) error {
	tx, _ := r.db.Beginx()
	var err error
	defer func() {
		switch r := recover(); {
		case r != recover():
			tx.Rollback()
			panic(r)
		case err != nil:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("INSERT INTO account (username, password_hash) VALUES (?, ?)", &account.Username, &account.PasswordHash); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *account) GetAccount(account *object.Account) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowx("SELECT * from account WHERE username = ?", &account.Username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}
	// db.closeする？
	return entity, nil
}
