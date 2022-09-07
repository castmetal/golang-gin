package common

import (
	"context"

	"github.com/google/uuid"
)

type IUseCase interface {
	Execute(ctx context.Context, dto IDTO) (bool, error)
}

type IDTO interface {
	Validate() (bool, error)
	ToBytes() ([]byte, error)
}

type IError interface {
	error
}

type IEntity interface {
	SetId(id uuid.UUID) *EntityBase
	GetId(entity *EntityBase) uuid.UUID
	GetEntity() *EntityBase
}

type IDatabase interface {
}

type IAggregateRoot interface {
}
