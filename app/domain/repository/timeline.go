package repository

import (
	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	GetPublicTimeline(timeline *object.Timeline) ([]*object.Status, error)
}
