package dao

import (
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	timeline struct {
		db *sqlx.DB
	}
)

func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

func (t *timeline) GetPublicTimeline(timeline *object.Timeline) ([]*object.Status, error) {
	var entities []*object.Status

	// if timeline.MaxID != nil{

	// }
	// if timeline.SinceID != nil{

	// }
	// if timeline.Limit != nil{

	// }
	rows, err := t.db.Queryx("SELECT * FROM status")
	if err != nil {
		return nil, fmt.Errorf("failed to select statuses : %w", err)
	}

	for rows.Next() {
		var status object.Status
		if err := rows.StructScan(&status); err != nil {
			return nil, fmt.Errorf("cannot mapping : %w", err)
		}
		entities = append(entities, &status)
	}

	return entities, nil
}
