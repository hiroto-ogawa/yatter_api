package repository

import (
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	CreateStatus(status *object.Status) error
	GetStatus(id int64) (*object.Status, error)
}
