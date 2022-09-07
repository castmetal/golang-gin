package common

import (
	"fmt"
	"net/http"
	"strings"
)

const PREFIX = "Domain - "
const (
	isNullOrEmptyText     string = PREFIX + "%s is null or empty !"
	invalidUserIdText     string = PREFIX + "The User '%s' is invalid!"
	invalidMessageText    string = PREFIX + "Message '%s' is invalid!"
	invalidConnectionText string = PREFIX + "Invalid Connection!"
	connectionClosedText  string = PREFIX + "Connection Closed!"
	onlySameMethodText    string = "Method not allowed"
	paramsErrorText       string = "Invalid Request"
	defaultDomainError    string = "System error, try again later"
	notFoundText          string = "%s not found"
	alreadyExistsText     string = "%s already exists"
)

var CodeErrors map[string]string = map[string]string{
	"INVALID_USER_ERROR":   "4001",
	"NULL_OR_EMPTY":        "4002",
	"INVALID_MESSAGE":      "4002",
	"INVALID_PARAMS":       "4003",
	"NOT_FOUND":            "4004",
	"ALREADY_EXISTS":       "4005",
	"DEFAULT_DOMAIN_ERROR": "5000",
	"INVALID_CONNECTION":   "5001",
	"CONNECTION_CLOSED":    "5002",
	"ONLY_SAME_METHOD":     "5003",
}

type ApplicationError struct {
	HttpError        int
	Code             string
	Msg              string
	ErrorDescription string
}

func (e *ApplicationError) Error() string {
	return e.Msg
}

func IsNullOrEmptyError(name string) *ApplicationError {
	return &ApplicationError{
		Msg:              fmt.Sprintf(isNullOrEmptyText, name),
		Code:             CodeErrors["NULL_OR_EMPTY"],
		HttpError:        http.StatusNotAcceptable,
		ErrorDescription: fmt.Sprintf(isNullOrEmptyText, name),
	}
}

func InvalidUserIdError(name string) *ApplicationError {
	return &ApplicationError{
		Msg:              fmt.Sprintf(invalidUserIdText, name),
		Code:             CodeErrors["INVALID_USER_ERROR"],
		HttpError:        http.StatusNotAcceptable,
		ErrorDescription: fmt.Sprintf(isNullOrEmptyText, name),
	}
}

func InvalidMessageError(name string) *ApplicationError {
	return &ApplicationError{
		Msg:              fmt.Sprintf(invalidMessageText, name),
		Code:             CodeErrors["INVALID_MESSAGE"],
		HttpError:        http.StatusNotAcceptable,
		ErrorDescription: fmt.Sprintf(isNullOrEmptyText, name),
	}
}

func InvalidConnectionError(description string) *ApplicationError {
	return &ApplicationError{
		Msg:              invalidConnectionText,
		Code:             CodeErrors["INVALID_CONNECTION"],
		HttpError:        http.StatusInternalServerError,
		ErrorDescription: description,
	}
}

func ConnectionClosedError(description string) *ApplicationError {
	return &ApplicationError{
		Msg:              connectionClosedText,
		Code:             CodeErrors["CONNECTION_CLOSED"],
		HttpError:        http.StatusInternalServerError,
		ErrorDescription: description,
	}
}

func DefaultDomainError(description string) *ApplicationError {
	return &ApplicationError{
		Msg:              defaultDomainError,
		Code:             CodeErrors["DEFAULT_DOMAIN_ERROR"],
		HttpError:        http.StatusInternalServerError,
		ErrorDescription: description,
	}
}

func OnlySameMethodError() *ApplicationError {
	return &ApplicationError{
		Msg:              onlySameMethodText,
		Code:             CodeErrors["ONLY_SAME_METHOD"],
		HttpError:        http.StatusForbidden,
		ErrorDescription: onlySameMethodText,
	}
}

func NotFoundError(name string) *ApplicationError {
	return &ApplicationError{
		Msg:              fmt.Sprintf(notFoundText, name),
		Code:             CodeErrors["NOT_FOUND"],
		HttpError:        http.StatusNotFound,
		ErrorDescription: fmt.Sprintf(isNullOrEmptyText, name),
	}
}

func InvalidParamsError(errorString string) *ApplicationError {
	var errorDescription string
	var itemSeparator string

	splitStringError := strings.Split(errorString, "\n")

	for _, itemError := range splitStringError {
		splitItemString := strings.Split(itemError, "Error:")
		errorDescription += itemSeparator + strings.Replace(splitItemString[1], "\n", "", -1)
		itemSeparator = "|"
	}

	return &ApplicationError{
		Msg:              "Error on content in these fields:",
		Code:             CodeErrors["INVALID_PARAMS"],
		HttpError:        http.StatusNotAcceptable,
		ErrorDescription: errorDescription,
	}
}

func AlreadyExistsError(name string) *ApplicationError {
	return &ApplicationError{
		Msg:              fmt.Sprintf(alreadyExistsText, name),
		Code:             CodeErrors["ALREADY_EXISTS"],
		HttpError:        http.StatusForbidden,
		ErrorDescription: fmt.Sprintf(isNullOrEmptyText, name),
	}
}
