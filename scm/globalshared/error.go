package globalshared

import (
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type (
	// ErrorInternal type
	ErrorInternal struct {
		Message string
	}

	// ErrorUsecase type
	ErrorUsecase struct {
		Message string
	}

	// ErrorNotFound type
	ErrorNotFound struct {
		Message string
	}
)

// NewErrorInternal type
func NewErrorInternal(msg string) *ErrorInternal {
	return &ErrorInternal{msg}
}

func (e *ErrorInternal) Error() string {
	return e.Message
}

// NewErrorUsecase type
func NewErrorUsecase(msg string) *ErrorUsecase {
	return &ErrorUsecase{msg}
}

func (e *ErrorUsecase) Error() string {
	return e.Message
}

// NewErrorNotFound type
func NewErrorNotFound(msg string) *ErrorNotFound {
	return &ErrorNotFound{msg}
}

func (e *ErrorNotFound) Error() string {
	return e.Message
}

// NewDBError type
func NewDBError(err error) error {
	if err == nil {
		return nil
	}
	if err == gorm.ErrRecordNotFound {
		return NewErrorNotFound(err.Error())
	}
	if strings.Contains(err.Error(), "duplicate") {
		return NewErrorUsecase(err.Error())
	}
	return NewErrorInternal(err.Error())
}

// GetErrorResponse message and code
func GetErrorResponse(err error) (int, string) {
	if _, ok := err.(*ErrorNotFound); ok {
		return http.StatusBadRequest, err.Error()
	}
	if _, ok := err.(*ErrorUsecase); ok {
		return http.StatusBadRequest, err.Error()
	}
	return http.StatusInternalServerError, "internal server error"
}

// GetErrorResponseDetail message and code
func GetErrorResponseDetail(err error) (int, string) {
	if _, ok := err.(*ErrorNotFound); ok {
		return http.StatusNotFound, err.Error()
	}
	return GetErrorResponse(err)
}

// ParseErrorContext type
func ParseErrorContext(err error) error {
	if err != nil && err.Error() == "context canceled" {
		return nil
	}
	return err
}
