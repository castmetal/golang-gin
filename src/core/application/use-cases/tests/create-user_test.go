package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"testing"

	_dtos "golang-gin/src/core/application/dtos"
	_use_cases "golang-gin/src/core/application/use-cases"
	_in_memory_tests "golang-gin/src/core/application/use-cases/tests/in-memory-tests"
)

type CreateUserTestStruct struct {
	message []byte
	expect  string
}

var TestDataCreateUser = []CreateUserTestStruct{
	{[]byte(`{"first_name":"Castmetal","last_name":"Metal","user_name":"castmetal","email":"email@gmail.com","password":"password1"}`), "email@gmail.com"},
	{[]byte(`{"first_name":"Castmetal","last_name":"Metal","user_name":"castmetal","email":"email@gmail.com","password":"password"}`), ""},
	{[]byte(`{"first_name":"Castmetal","last_name":"Metal","user_name":"castmetal","email":"wrongemail","password":"password"}`), ""},
}

// Testing CreateUser Use Case
func TestCreateUser(t *testing.T) {
	var message io.Reader
	userRepository := _in_memory_tests.NewUserRepositoryFromConfig()
	_in_memory_tests.ClearDataUser()

	useCase, _ := _use_cases.NewCreateUser(userRepository)

	for _, testItem := range TestDataCreateUser {
		message = bytes.NewReader(testItem.message)
		var createUserDTO _dtos.CreateUserDTO = _dtos.CreateUserDTO{}

		messageBuffer := &bytes.Buffer{}
		messageBuffer.ReadFrom(message)

		err := json.Unmarshal(messageBuffer.Bytes(), &createUserDTO)

		res, err := useCase.Execute(context.Background(), createUserDTO)
		if res.User.Email != testItem.expect {

			t.Error(err.Error())
		}
	}

}
