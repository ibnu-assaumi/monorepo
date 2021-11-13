package globalshared

import (
	"net/http"
)

type ErrorDB struct {
	Message string
}

func (e *ErrorDB) Error() string {
	return e.Message
}

func NewErrorDB(msg string) *ErrorDB {
	return &ErrorDB{Message: msg}
}

func GetErrorResponse(err error) (int, string) {
	if _, ok := err.(*ErrorDB); ok {
		return http.StatusInternalServerError, "internal server error"
	}
	return http.StatusBadRequest, err.Error()
}
