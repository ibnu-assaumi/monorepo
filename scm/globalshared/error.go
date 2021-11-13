package globalshared

import (
	"net/http"
	"strings"

	"gorm.io/gorm"
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

func NewErrorDBGorm(err error) error {
	if err != nil && err != gorm.ErrRecordNotFound && !strings.Contains(err.Error(), "context") {
		return NewErrorDB(err.Error())
	}
	return err
}

func GetErrorResponse(err error) (int, string) {
	if _, ok := err.(*ErrorDB); ok {
		return http.StatusInternalServerError, "internal server error"
	}
	return http.StatusBadRequest, err.Error()
}
