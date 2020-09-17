package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RestError interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restError struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (error restError) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		error.ErrMessage, error.ErrStatus, error.ErrError, error.ErrCauses)
}

func (error restError) Message() string {
	return error.ErrMessage
}

func (error restError) Status() int {
	return error.ErrStatus
}

func (error restError) Causes() []interface{} {
	return error.ErrCauses
}

func NewRestError(message string, status int, err string, causes []interface{}) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

func NewRestErrorFromBytes(bytes []byte) (RestError, error) {
	var restErr restError
	if err := json.Unmarshal(bytes, &restErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return restErr, nil
}

func NewBadRequestError(message string) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NewNotFoundError(message string) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

func NewUnauthorizedError(message string) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

func NewInternalServerError(message string, err error) RestError {
	returnError := restError{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		returnError.ErrCauses = append(returnError.ErrCauses, err.Error())
	}
	return returnError
}
