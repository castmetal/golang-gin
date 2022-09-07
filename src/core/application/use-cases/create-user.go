package use_cases

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"

	_dtos "golang-gin/src/core/application/dtos"
	_common "golang-gin/src/core/domains/common"
	_user "golang-gin/src/core/domains/user"
	_redis "golang-gin/src/infra/redis"
)

type (
	CreateUser interface {
		Execute(ctx context.Context, createUserDTO _dtos.CreateUserDTO) (_dtos.CreateUserResponseDTO, error)
		toResponse(user *_user.User) _dtos.CreateUserResponseDTO
	}
	CreateUserRequest struct {
		CreateUser
		Repository  _user.IUserRepository
		RedisClient _redis.IRedisClient
	}
)

func NewCreateUser(repository _user.IUserRepository, redisClient _redis.IRedisClient) (CreateUser, error) {
	var uc CreateUser = &CreateUserRequest{
		Repository:  repository,
		RedisClient: redisClient,
	}

	return uc, nil
}

// Put here your validation message and return your struct mapper to service
func (uc *CreateUserRequest) Execute(ctx context.Context, createUserDTO _dtos.CreateUserDTO) (_dtos.CreateUserResponseDTO, error) {
	var response = _dtos.CreateUserResponseDTO{}

	_, err := createUserDTO.Validate()
	if err != nil {
		return response, _common.InvalidParamsError(err.Error())
	}

	dtoBytes, err := createUserDTO.ToBytes()
	if err != nil {
		return response, _common.DefaultDomainError(err.Error())
	}

	var dtoReader io.Reader = bytes.NewReader(dtoBytes)

	var userProps = getUserProps(dtoReader)

	user, err := _user.NewUserEntity(userProps)
	if err != nil {
		return response, _common.DefaultDomainError(err.Error())
	}

	res, err := uc.Repository.FindOneByEmail(ctx, user.Email)
	if res != nil && res.Email != "" {
		return response, _common.AlreadyExistsError(user.Email)
	}

	_, err = uc.Repository.Create(ctx, user)
	if err != nil {
		return response, _common.DefaultDomainError(err.Error())
	}

	go func(ctx context.Context) {
		_ = uc.RedisClient.DelData(ctx, _user.REDIS_LIST_ALL_USERS_KEY)
	}(ctx)

	return uc.toResponse(user), nil
}

func (uc *CreateUserRequest) toResponse(user *_user.User) _dtos.CreateUserResponseDTO {
	return _dtos.CreateUserResponseDTO{
		User: _dtos.UserResponseDTO{
			ID:        user.ID.String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			UserName:  user.UserName,
		},
	}
}

func getUserProps(message io.Reader) _user.UserProps {
	var userProps _user.UserProps
	messageBuffer := &bytes.Buffer{}
	messageBuffer.ReadFrom(message)

	if err := json.Unmarshal(messageBuffer.Bytes(), &userProps); err != nil {
		log.Fatal(err)
	}

	return userProps
}
