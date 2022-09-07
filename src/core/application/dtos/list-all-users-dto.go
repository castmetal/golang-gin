package dtos

import (
	"bytes"
	"encoding/json"
	_common "golang-gin/src/core/domains/common"
	"io"

	"github.com/go-playground/validator/v10"
)

type (
	IListAllUsersDTO interface {
		_common.IDTO
		New(message io.Reader) (IListAllUsersDTO, error)
	}

	ListAllUsersDTO struct {
		IListAllUsersDTO
		Limit  int `json:"limit" validate:"numeric,min=0,max=100"`
		Offset int `json:"offset" validate:"numeric,min=0,max=10000000"`
	}
)

func (dto *ListAllUsersDTO) New(message io.Reader) (IListAllUsersDTO, error) {
	var IDTO IListAllUsersDTO = &ListAllUsersDTO{}

	messageBuffer := &bytes.Buffer{}
	messageBuffer.ReadFrom(message)

	err := json.Unmarshal(messageBuffer.Bytes(), &IDTO)
	if err != nil {
		return IDTO, err
	}

	return IDTO, nil
}

func (dto *ListAllUsersDTO) Validate() (bool, error) {
	var validate *validator.Validate

	validate = validator.New()
	err := validate.Struct(dto)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dto *ListAllUsersDTO) ToBytes() ([]byte, error) {
	b, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	return b, nil
}
