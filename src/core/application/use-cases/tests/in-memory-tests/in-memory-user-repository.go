package in_memory_tests

import (
	"context"

	_common "golang-gin/src/core/domains/common"
	_user "golang-gin/src/core/domains/user"
)

const _collectionName = "UsersInMemory"

var dbDataUser map[string]_user.User = make(map[string]_user.User)

type userRepository struct {
	db _common.IDatabase
}

func NewUserRepositoryFromConfig() _user.IUserRepository {
	var db _common.IDatabase

	return newUserRepository(db)
}

func newUserRepository(db _common.IDatabase) _user.IUserRepository {
	return &userRepository{db: db}
}

func ClearDataUser() {
	dbDataUser = make(map[string]_user.User)
}

func (repository userRepository) FindOneById(ctx context.Context, id string) (*_user.User, error) {
	var user _user.User = dbDataUser[id]

	userId := string(user.ID[:])

	if userId == "" {
		return nil, _common.NotFoundError("User")
	}

	return &user, nil
}

func (repository userRepository) Create(ctx context.Context, user *_user.User) (*_user.User, error) {
	var u *_user.User

	u, _ = repository.FindOneByEmail(ctx, user.Email)
	if u != nil {
		return nil, _common.AlreadyExistsError("User")
	}

	userId := string(user.ID[:])
	dbDataUser[userId] = *user

	return user, nil
}

func (repository userRepository) FindOneByEmail(ctx context.Context, email string) (*_user.User, error) {
	var user *_user.User

	for _, u := range dbDataUser {
		if u.Email == email {
			user = &u
		}
	}

	if user == nil {
		return nil, _common.NotFoundError("User")
	}

	return user, nil
}

func (repository userRepository) ListAll(ctx context.Context, limit int, offset int) ([]*_user.User, error) {
	var currentOffet int = 0
	var numItems int = 0
	var response []*_user.User = []*_user.User{}

	for _, u := range dbDataUser {
		if currentOffet == offset {
			response = append(response, &u)
		}

		if numItems+1 == limit {
			currentOffet++
			numItems = 0
			continue
		}

		numItems++
	}

	return response, nil
}

func (repository userRepository) GetTotalRows(ctx context.Context) int64 {
	var totalRows int64 = int64(len(dbDataUser))

	return totalRows
}
