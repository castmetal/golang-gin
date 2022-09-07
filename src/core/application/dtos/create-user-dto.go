package dtos

import (
	"bytes"
	"encoding/json"
	_common "golang-gin/src/core/domains/common"
	"io"

	"github.com/go-playground/validator/v10"
)

type (
	ICreateUserDTO interface {
		_common.IDTO
		New(message io.Reader) (ICreateUserDTO, error)
	}

	CreateUserDTO struct {
		ICreateUserDTO
		UserDTO
	}
)

func (dto *CreateUserDTO) New(message io.Reader) (ICreateUserDTO, error) {
	var IDTO ICreateUserDTO = &CreateUserDTO{}

	messageBuffer := &bytes.Buffer{}
	messageBuffer.ReadFrom(message)

	err := json.Unmarshal(messageBuffer.Bytes(), &IDTO)
	if err != nil {
		return IDTO, err
	}

	return IDTO, nil
}

func (dto *CreateUserDTO) Validate() (bool, error) {
	var validate *validator.Validate

	validate = validator.New()
	err := validate.Struct(dto)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dto *CreateUserDTO) ToBytes() ([]byte, error) {
	b, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	return b, nil
}
