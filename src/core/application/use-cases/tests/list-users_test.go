package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"testing"

	_dtos "golang-gin/src/core/application/dtos"
	_use_cases "golang-gin/src/core/application/use-cases"
	_in_memory_tests "golang-gin/src/core/application/use-cases/tests/in-memory-tests"
)

type ListAllUsersTestStruct struct {
	message []byte
	expect  int
}

var TestDataListAllUsers = []ListAllUsersTestStruct{
	{[]byte(`{"limit":10,"offset":0}`), 0},
	{[]byte(`{"limit":10,"offset":0}`), 0},
}

// Testing ListAllUsers Use Case
func TestListAllUsers(t *testing.T) {
	var message io.Reader
	userRepository := _in_memory_tests.NewUserRepositoryFromConfig()
	_in_memory_tests.ClearDataUser()

	useCase, _ := _use_cases.NewListAllUsers(userRepository)

	for _, testItem := range TestDataListAllUsers {
		message = bytes.NewReader(testItem.message)
		var listAllUsersDTO _dtos.ListAllUsersDTO = _dtos.ListAllUsersDTO{
			Limit:  0,
			Offset: 10,
		}

		messageBuffer := &bytes.Buffer{}
		messageBuffer.ReadFrom(message)

		err := json.Unmarshal(messageBuffer.Bytes(), &listAllUsersDTO)
		if err != nil {

			t.Error(err.Error())
		}

		res, err := useCase.Execute(context.Background(), listAllUsersDTO)
		if res.Users == nil || len(res.Users) != testItem.expect {
			fmt.Println("Erro", err.Error())
			t.Error(err.Error())
		}
	}

}
