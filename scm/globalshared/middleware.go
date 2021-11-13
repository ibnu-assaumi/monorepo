package globalshared

import (
	"fmt"
	"net/http"

	"github.com/Bhinneka/candi/tracer"
	"github.com/Bhinneka/candi/wrapper"
	"github.com/labstack/echo"
)

// HTTPPanicMiddleware echo middleware
func HTTPPanicMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					ctx := c.Request().Context()
					err := fmt.Errorf("PANIC: %v", r)
					tracer.SetError(ctx, err)
					Log(ctx, LogParam{
						Error:         err,
						Message:       err.Error(),
						OperationName: c.Request().Host + c.Request().RequestURI,
						Scope:         c.Request().Method,
						IsSentry:      true,
					})
					SlackSend(ctx, SlackParam{
						Title:         c.Request().Method + " " + c.Request().Host + c.Request().RequestURI,
						OperationName: "REST Request",
						Error:         err,
					})
					wrapper.NewHTTPResponse(http.StatusInternalServerError, "internal server error").JSON(c.Response())
				}
			}()

			return next(c)
		}
	}
}
