package globalshared

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo"

	"github.com/Bhinneka/candi/candihelper"
	"github.com/Bhinneka/candi/candishared"
	"github.com/Bhinneka/candi/codebase/interfaces"
)

const invalidParameter = "invalid parameter"

// EchoValidateQueryParam : validate echo request url query parameter
func EchoValidateQueryParam(c echo.Context, param interface{}, schemaID string, validator interfaces.Validator) error {
	multiError := candihelper.NewMultiError()
	if err := candihelper.ParseFromQueryParam(c.Request().URL.Query(), param); err != nil {
		multiError.Append(invalidParameter, err)
		return multiError
	}
	b, err := json.Marshal(param)
	if err != nil {
		multiError.Append(invalidParameter, err)
		return multiError
	}
	if err := validator.ValidateDocument(schemaID, b); err != nil {
		multiError.Append(invalidParameter, err)
		return multiError
	}
	return nil
}

// EchoValidateBodyParam : validate echo request body parameter
func EchoValidateBodyParam(c echo.Context, param interface{}, schemaID string, validator interfaces.Validator) error {
	var (
		multiError = candihelper.NewMultiError()
		bodyBytes  []byte
		err        error
	)
	if c.Request().Body != nil {
		bodyBytes, err = ioutil.ReadAll(c.Request().Body)
		if err != nil {
			multiError.Append(invalidParameter, err)
			return multiError
		}
	}

	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if err = validator.ValidateDocument(schemaID, bodyBytes); err != nil {
		multiError.Append(invalidParameter, err)
		return multiError
	}

	if err := json.Unmarshal(bodyBytes, param); err != nil {
		multiError.Append(invalidParameter, err)
		return multiError
	}

	return nil
}

func EchoGetPageLimit(filter candishared.Filter, totalData int) (int, int) {
	var (
		page  = filter.Page
		limit = filter.Limit
	)
	if filter.ShowAll {
		page = 1
		limit = int(totalData)
	}

	return page, limit

}
