package dtos

import (
	"bytes"
	"encoding/json"
	_common "golang-gin/src/core/domains/common"
	"io"

	"github.com/go-playground/validator/v10"
)

type (
	IListAllUsersResponseDTO interface {
		_common.IDTO
		New(message io.Reader) (IListAllUsersResponseDTO, error)
	}

	ListAllUsersResponseDTO struct {
		Users      []UserResponseDTO `json:"user"`
		TotalRows  int64             `json:"total_rows"`
		TotalPages int               `json:"total_pages"`
		Limit      int               `json:"limit"`
		Offset     int               `json:"page"`
	}
)

func (dto *ListAllUsersResponseDTO) New(message io.Reader) (IListAllUsersResponseDTO, error) {
	var IDTO IListAllUsersResponseDTO = &ListAllUsersResponseDTO{}

	messageBuffer := &bytes.Buffer{}
	messageBuffer.ReadFrom(message)

	err := json.Unmarshal(messageBuffer.Bytes(), &IDTO)
	if err != nil {
		return IDTO, err
	}

	return IDTO, nil
}

func (dto *ListAllUsersResponseDTO) Validate() (bool, error) {
	var validate *validator.Validate

	validate = validator.New()
	err := validate.Struct(dto)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dto *ListAllUsersResponseDTO) ToBytes() ([]byte, error) {
	b, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	return b, nil
}
