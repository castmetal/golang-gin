package common

import (
	"net/http"

	"github.com/google/uuid"
)

type EntityBase struct {
	ID uuid.UUID `json:"id" bson:"_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
}

func NewAbstractEntity(idString string) *EntityBase {
	var _id uuid.UUID

	if idString != "" {
		_id = uuid.Must(uuid.FromBytes([]byte(idString)))
	} else {
		_id = uuid.New()
	}

	var entity = &EntityBase{
		ID: _id,
	}

	return entity
}

type Controller struct {
	URL    string
	Method string
	Handle func(w http.ResponseWriter, r *http.Request)
}
