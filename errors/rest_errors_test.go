package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewRestError(t *testing.T) {
	// TODO
}

func TestNewRestErrorFromBytes(t *testing.T) {
	// TODO
}

func TestNewBadRequestError(t *testing.T) {
	// TODO
}

func TestNewNotFoundError(t *testing.T) {
	// TODO
}

func TestNewUnauthorizedError(t *testing.T) {
	// TODO
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("This is an error message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "This is an error message", err.Message())
	assert.EqualValues(t, "message: This is an error message - status: 500 - error: internal_server_error - causes: [database error]", err.Error())

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])
}
