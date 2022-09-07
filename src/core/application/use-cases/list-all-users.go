package use_cases

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"math"
	"time"

	_dtos "golang-gin/src/core/application/dtos"
	_common "golang-gin/src/core/domains/common"
	_user "golang-gin/src/core/domains/user"
	_redis "golang-gin/src/infra/redis"

	"github.com/goccy/go-json"
)

type (
	ListUsers interface {
		Execute(ctx context.Context, listAllUsersDTO _dtos.ListAllUsersDTO) (_dtos.ListAllUsersResponseDTO, error)
		toResponse(users []*_user.User, totalRows int64, listAllUsersDTO _dtos.ListAllUsersDTO) _dtos.ListAllUsersResponseDTO
	}
	ListAllUsersRequest struct {
		ListUsers
		Repository  _user.IUserRepository
		RedisClient _redis.IRedisClient
	}
)

func NewListAllUsers(repository _user.IUserRepository, redisClient _redis.IRedisClient) (ListUsers, error) {
	var uc ListUsers = &ListAllUsersRequest{
		Repository:  repository,
		RedisClient: redisClient,
	}

	return uc, nil
}

// Put here your validation message and return your struct mapper to service
func (uc *ListAllUsersRequest) Execute(ctx context.Context, listAllUsersDTO _dtos.ListAllUsersDTO) (_dtos.ListAllUsersResponseDTO, error) {
	var response = _dtos.ListAllUsersResponseDTO{}
	var cacheDuration = 48 * time.Hour

	_, err := listAllUsersDTO.Validate()
	if err != nil {
		return response, _common.InvalidParamsError(err.Error())
	}

	paramsBytes, err := json.Marshal(listAllUsersDTO)
	if err != nil {
		return response, _common.DefaultDomainError(err.Error())
	}

	var cacheHash = md5.Sum([]byte(string(paramsBytes)))
	cacheKey := _user.REDIS_LIST_ALL_USERS_KEY + hex.EncodeToString(cacheHash[:])

	getCache, err := uc.RedisClient.GetData(ctx, cacheKey)
	if err == nil && getCache != "" {
		json.Unmarshal([]byte(getCache), &response)

		return response, nil
	}

	listData, err := uc.Repository.ListAll(ctx, listAllUsersDTO.Limit, listAllUsersDTO.Offset)
	if err != nil {
		return response, _common.DefaultDomainError(err.Error())
	}

	totalRows := uc.Repository.GetTotalRows(ctx)

	resp := uc.toResponse(listData, totalRows, listAllUsersDTO)
	responseBytes, err := json.Marshal(resp)
	if err != nil {
		return response, _common.DefaultDomainError(err.Error())
	}

	go func(ctx context.Context, cacheKey string, responseBytes string, cacheDuration time.Duration) {
		_ = uc.RedisClient.SetData(ctx, cacheKey, responseBytes, cacheDuration)
	}(ctx, cacheKey, string(responseBytes), cacheDuration)

	return resp, nil
}

func (uc *ListAllUsersRequest) toResponse(users []*_user.User, totalRows int64, listAllUsersDTO _dtos.ListAllUsersDTO) _dtos.ListAllUsersResponseDTO {
	var usersData []_dtos.UserResponseDTO = []_dtos.UserResponseDTO{}

	for _, item := range users {
		usersData = append(usersData, _dtos.UserResponseDTO{
			ID:        item.ID.String(),
			FirstName: item.FirstName,
			LastName:  item.LastName,
			Email:     item.Email,
			UserName:  item.UserName,
		})
	}

	return _dtos.ListAllUsersResponseDTO{
		Users:      usersData,
		Limit:      listAllUsersDTO.Limit,
		Offset:     listAllUsersDTO.Offset,
		TotalRows:  totalRows,
		TotalPages: int(math.Ceil(float64(totalRows) / float64(listAllUsersDTO.Limit))),
	}
}
