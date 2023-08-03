package repository

import (
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	CreateStatus(status *object.Status) error
}
