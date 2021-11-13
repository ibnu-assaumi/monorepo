// Code generated by candi v1.8.8.

package resthandler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mockusecase "monorepo/services/seaotter/pkg/mocks/modules/salesorder/usecase"
	mocksharedusecase "monorepo/services/seaotter/pkg/mocks/shared/usecase"
	shareddomain "monorepo/services/seaotter/pkg/shared/domain"

	mockdeps "/mocks/codebase/factory/dependency"
	mockinterfaces "/mocks/codebase/interfaces"

	"github.com/Bhinneka/candi/candishared"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name, reqBody                       string
	wantValidateError, wantUsecaseError error
	wantRespCode                        int
}

var (
	errFoo = errors.New("Something error")
)

func TestNewRestHandler(t *testing.T) {
	mockMiddleware := &mockinterfaces.Middleware{}
	mockMiddleware.On("HTTPPermissionACL", mock.Anything).Return(func(http.Handler) http.Handler { return nil })
	mockValidator := &mockinterfaces.Validator{}

	mockDeps := &mockdeps.Dependency{}
	mockDeps.On("GetMiddleware").Return(mockMiddleware)
	mockDeps.On("GetValidator").Return(mockValidator)

	handler := NewRestHandler(nil, mockDeps)
	assert.NotNil(t, handler)

	e := echo.New()
	handler.Mount(e.Group("/"))
}

func TestRestHandler_getAllSalesorder(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", wantUsecaseError: nil, wantRespCode: 200,
		},
		{
			name: "Testcase #2: Negative", reqBody: "?page=str", wantUsecaseError: errFoo, wantRespCode: 400,
		},
		{
			name: "Testcase #3: Negative", wantUsecaseError: errFoo, wantRespCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			salesorderUsecase := &mockusecase.SalesorderUsecase{}
			salesorderUsecase.On("GetAllSalesorder", mock.Anything, mock.Anything).Return(
				[]shareddomain.Salesorder{}, candishared.Meta{}, tt.wantUsecaseError)
			mockValidator := &mockinterfaces.Validator{}
			mockValidator.On("ValidateDocument", mock.Anything, mock.Anything).Return(tt.wantValidateError)

			uc := &mocksharedusecase.Usecase{}
			uc.On("Salesorder").Return(salesorderUsecase)

			handler := RestHandler{uc: uc, validator: mockValidator}

			req := httptest.NewRequest(http.MethodGet, "/"+tt.reqBody, strings.NewReader(tt.reqBody))
			req = req.WithContext(candishared.SetToContext(req.Context(), candishared.ContextKeyTokenClaim, &candishared.TokenClaim{}))
			req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()
			echoContext := echo.New().NewContext(req, res)
			err := handler.getAllSalesorder(echoContext)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantRespCode, res.Code)
		})
	}
}

func TestRestHandler_getDetailSalesorderByID(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", wantUsecaseError: nil, wantRespCode: 200,
		},
		{
			name: "Testcase #2: Negative", wantUsecaseError: errFoo, wantRespCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			salesorderUsecase := &mockusecase.SalesorderUsecase{}
			salesorderUsecase.On("GetDetailSalesorder", mock.Anything, mock.Anything).Return(shareddomain.Salesorder{}, tt.wantUsecaseError)
			mockValidator := &mockinterfaces.Validator{}
			mockValidator.On("ValidateDocument", mock.Anything, mock.Anything).Return(tt.wantValidateError)

			uc := &mocksharedusecase.Usecase{}
			uc.On("Salesorder").Return(salesorderUsecase)

			handler := RestHandler{uc: uc, validator: mockValidator}

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.reqBody))
			req = req.WithContext(candishared.SetToContext(req.Context(), candishared.ContextKeyTokenClaim, &candishared.TokenClaim{}))
			req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()
			echoContext := echo.New().NewContext(req, res)
			err := handler.getDetailSalesorderByID(echoContext)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantRespCode, res.Code)
		})
	}
}

func TestRestHandler_createSalesorder(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", reqBody: `{"email": "test@test.com"}`, wantUsecaseError: nil, wantRespCode: 200,
		},
		{
			name: "Testcase #2: Negative", reqBody: `{"email": test@test.com}`, wantUsecaseError: nil, wantRespCode: 400,
		},
		{
			name: "Testcase #3: Negative", reqBody: `{"email": "test@test.com"}`, wantUsecaseError: errFoo, wantRespCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			salesorderUsecase := &mockusecase.SalesorderUsecase{}
			salesorderUsecase.On("CreateSalesorder", mock.Anything, mock.Anything).Return(tt.wantUsecaseError)
			mockValidator := &mockinterfaces.Validator{}
			mockValidator.On("ValidateDocument", mock.Anything, mock.Anything).Return(tt.wantValidateError)

			uc := &mocksharedusecase.Usecase{}
			uc.On("Salesorder").Return(salesorderUsecase)

			handler := RestHandler{uc: uc, validator: mockValidator}

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.reqBody))
			req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()
			echoContext := echo.New().NewContext(req, res)
			echoContext.SetParamNames("id")
			echoContext.SetParamValues("001")
			err := handler.createSalesorder(echoContext)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantRespCode, res.Code)
		})
	}
}

func TestRestHandler_updateSalesorder(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", reqBody: `{"email": "test@test.com"}`, wantUsecaseError: nil, wantRespCode: 200,
		},
		{
			name: "Testcase #2: Negative", reqBody: `{"email": test@test.com}`, wantValidateError: errFoo, wantRespCode: 400,
		},
		{
			name: "Testcase #3: Negative", reqBody: `{"email": test@test.com}`, wantUsecaseError: nil, wantRespCode: 400,
		},
		{
			name: "Testcase #4: Negative", reqBody: `{"email": "test@test.com"}`, wantUsecaseError: errFoo, wantRespCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			salesorderUsecase := &mockusecase.SalesorderUsecase{}
			salesorderUsecase.On("UpdateSalesorder", mock.Anything, mock.Anything, mock.Anything).Return(tt.wantUsecaseError)
			mockValidator := &mockinterfaces.Validator{}
			mockValidator.On("ValidateDocument", mock.Anything, mock.Anything).Return(tt.wantValidateError)

			uc := &mocksharedusecase.Usecase{}
			uc.On("Salesorder").Return(salesorderUsecase)

			handler := RestHandler{uc: uc, validator: mockValidator}

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.reqBody))
			req = req.WithContext(candishared.SetToContext(req.Context(), candishared.ContextKeyTokenClaim, &candishared.TokenClaim{}))
			req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()
			echoContext := echo.New().NewContext(req, res)
			err := handler.updateSalesorder(echoContext)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantRespCode, res.Code)
		})
	}
}

func TestRestHandler_deleteSalesorder(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", wantUsecaseError: nil, wantRespCode: 200,
		},
		{
			name: "Testcase #2: Negative", wantUsecaseError: errFoo, wantRespCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			salesorderUsecase := &mockusecase.SalesorderUsecase{}
			salesorderUsecase.On("DeleteSalesorder", mock.Anything, mock.Anything).Return(tt.wantUsecaseError)
			mockValidator := &mockinterfaces.Validator{}
			mockValidator.On("ValidateDocument", mock.Anything, mock.Anything).Return(tt.wantValidateError)

			uc := &mocksharedusecase.Usecase{}
			uc.On("Salesorder").Return(salesorderUsecase)

			handler := RestHandler{uc: uc, validator: mockValidator}

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.reqBody))
			req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()
			echoContext := echo.New().NewContext(req, res)
			err := handler.deleteSalesorder(echoContext)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantRespCode, res.Code)
		})
	}
}