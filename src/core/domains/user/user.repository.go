package user

import (
	"context"
	_common "golang-gin/src/core/domains/common"
)

type IUserRepository interface {
	_common.IAggregateRoot
	Create(ctx context.Context, user *User) (*User, error)
	FindOneByEmail(ctx context.Context, email string) (*User, error)
	ListAll(ctx context.Context, limit int, offset int) ([]*User, error)
	GetTotalRows(ctx context.Context) int64
}
