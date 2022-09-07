package controllers_rest_v1

import (
	"bytes"
	"context"
	"encoding/json"
	_dtos "golang-gin/src/core/application/dtos"
	_use_cases "golang-gin/src/core/application/use-cases"
	_common "golang-gin/src/core/domains/common"
	_redis "golang-gin/src/infra/redis"
	_repositories "golang-gin/src/infra/repositories"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	VERSION = "/v1"
)

func CreateUserControllerV1(c *gin.Context) {
	var bodyBytes []byte

	if c.Request.Body != nil {
		defer c.Request.Body.Close()
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	var createUserDTO = &_dtos.CreateUserDTO{}
	err := json.Unmarshal(bodyBytes, &createUserDTO)
	if err != nil {
		errMessage := _common.InvalidParamsError(err.Error())

		_common.HandleHttpErrors(errMessage, c)
		return
	}

	userRepository := _repositories.NewUserRepositoryFromConfig()
	redisClient := _redis.NewRedisClient()

	createUser, err := _use_cases.NewCreateUser(userRepository, redisClient)
	if err != nil {
		_common.HandleHttpErrors(err, c)
		return
	}

	context := context.Background()

	result, err := createUser.Execute(context, *createUserDTO)
	if err != nil {
		_common.HandleHttpErrors(err, c)
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
	return
}

func ListAllUsersV1(c *gin.Context) {
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		limit = 10
	}

	var listAllUsersDTO = &_dtos.ListAllUsersDTO{
		Limit:  limit,
		Offset: offset,
	}

	userRepository := _repositories.NewUserRepositoryFromConfig()
	redisClient := _redis.NewRedisClient()

	createUser, err := _use_cases.NewListAllUsers(userRepository, redisClient)
	if err != nil {
		_common.HandleHttpErrors(err, c)
		return
	}

	context := context.Background()

	result, err := createUser.Execute(context, *listAllUsersDTO)
	if err != nil {
		_common.HandleHttpErrors(err, c)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
	return
}

func SetUserControllers(routerEngine *gin.Engine) {
	routerEngine.POST(VERSION+"/users/create", CreateUserControllerV1)
	routerEngine.GET(VERSION+"/users", ListAllUsersV1)
}
