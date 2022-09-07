package dtos

import (
	"bytes"
	"encoding/json"
	_common "golang-gin/src/core/domains/common"
	"io"

	"github.com/go-playground/validator/v10"
)

type (
	ICreateUserResponseDTO interface {
		_common.IDTO
		New(message io.Reader) (ICreateUserDTO, error)
	}

	UserResponseDTO struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name" validate:"required,min=2"`
		LastName  string `json:"last_name" validate:"required,min=2"`
		UserName  string `json:"user_name" validate:"required,min=2"`
		Email     string `json:"email" validate:"required,email"`
	}

	CreateUserResponseDTO struct {
		User UserResponseDTO `json:"user"`
	}
)

func (dto *CreateUserResponseDTO) New(message io.Reader) (ICreateUserResponseDTO, error) {
	var IDTO ICreateUserDTO = &CreateUserDTO{}

	messageBuffer := &bytes.Buffer{}
	messageBuffer.ReadFrom(message)

	err := json.Unmarshal(messageBuffer.Bytes(), &IDTO)
	if err != nil {
		return IDTO, err
	}

	return IDTO, nil
}

func (dto *CreateUserResponseDTO) Validate() (bool, error) {
	var validate *validator.Validate

	validate = validator.New()
	err := validate.Struct(dto)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dto *CreateUserResponseDTO) ToBytes() ([]byte, error) {
	b, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	return b, nil
}
