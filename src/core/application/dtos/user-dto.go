package dtos

import (
	"bytes"
	"encoding/json"
	"io"

	_common "golang-gin/src/core/domains/common"

	"github.com/go-playground/validator/v10"
)

type (
	IUserDTO interface {
		_common.IDTO
		New(message io.Reader) (IUserDTO, error)
	}

	UserDTO struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name" validate:"required,min=2"`
		LastName  string `json:"last_name" validate:"required,min=2"`
		UserName  string `json:"user_name" validate:"required,min=2"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required,min=8"`
	}
)

func (dto *UserDTO) New(message io.Reader) (IUserDTO, error) {
	var IDTO IUserDTO = &UserDTO{}

	messageBuffer := &bytes.Buffer{}
	messageBuffer.ReadFrom(message)

	err := json.Unmarshal(messageBuffer.Bytes(), &IDTO)
	if err != nil {
		return IDTO, err
	}

	return IDTO, nil
}

func (dto *UserDTO) Validate() (bool, error) {
	var validate *validator.Validate

	validate = validator.New()
	err := validate.Struct(dto)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dto *UserDTO) ToBytes() ([]byte, error) {
	b, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	return b, nil
}
