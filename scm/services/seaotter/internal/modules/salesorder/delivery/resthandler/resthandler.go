// Code generated by candi v1.8.8.

package resthandler

import (
	"net/http"

	"github.com/labstack/echo"

	"monorepo/services/seaotter/internal/modules/salesorder/domain"
	shareddomain "monorepo/services/seaotter/pkg/shared/domain"
	"monorepo/services/seaotter/pkg/shared/usecase"

	"github.com/Bhinneka/candi/candihelper"
	"github.com/Bhinneka/candi/codebase/factory/dependency"
	"github.com/Bhinneka/candi/codebase/interfaces"
	"github.com/Bhinneka/candi/tracer"
	"github.com/Bhinneka/candi/wrapper"

	"github.com/Bhinneka/candi/candishared"
)

// RestHandler handler
type RestHandler struct {
	mw        interfaces.Middleware
	uc        usecase.Usecase
	validator interfaces.Validator
}

// NewRestHandler create new rest handler
func NewRestHandler(uc usecase.Usecase, deps dependency.Dependency) *RestHandler {
	return &RestHandler{
		uc: uc, mw: deps.GetMiddleware(), validator: deps.GetValidator(),
	}
}

// Mount handler with root "/"
// handling version in here
func (h *RestHandler) Mount(root *echo.Group) {
	v1Root := root.Group(candihelper.V1)

	salesorder := v1Root.Group("/salesorder", echo.WrapMiddleware(h.mw.HTTPBearerAuth))
	salesorder.GET("", h.getAllSalesorder, echo.WrapMiddleware(h.mw.HTTPPermissionACL("resource.public")))
	salesorder.GET("/:id", h.getDetailSalesorderByID, echo.WrapMiddleware(h.mw.HTTPPermissionACL("resource.public")))
	salesorder.POST("", h.createSalesorder, echo.WrapMiddleware(h.mw.HTTPPermissionACL("resource.public")))
	salesorder.PUT("/:id", h.updateSalesorder, echo.WrapMiddleware(h.mw.HTTPPermissionACL("resource.public")))
	salesorder.DELETE("/:id", h.deleteSalesorder, echo.WrapMiddleware(h.mw.HTTPPermissionACL("resource.public")))
}

func (h *RestHandler) getAllSalesorder(c echo.Context) error {
	trace, ctx := tracer.StartTraceWithContext(c.Request().Context(), "SalesorderDeliveryREST:GetAllSalesorder")
	defer trace.Finish()

	tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using HTTPBearerAuth in middleware for this handler

	var filter domain.FilterSalesorder
	if err := candihelper.ParseFromQueryParam(c.Request().URL.Query(), &filter); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	data, meta, err := h.uc.Salesorder().GetAllSalesorder(ctx, &filter)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	message := "Success, with your user id (" + tokenClaim.Subject + ") and role (" + tokenClaim.Role + ")"
	return wrapper.NewHTTPResponse(http.StatusOK, message, meta, data).JSON(c.Response())
}

func (h *RestHandler) getDetailSalesorderByID(c echo.Context) error {
	trace, ctx := tracer.StartTraceWithContext(c.Request().Context(), "SalesorderDeliveryREST:GetDetailSalesorderByID")
	defer trace.Finish()

	data, err := h.uc.Salesorder().GetDetailSalesorder(ctx, c.Param("id"))
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "Success", data).JSON(c.Response())
}

func (h *RestHandler) createSalesorder(c echo.Context) error {
	trace, ctx := tracer.StartTraceWithContext(c.Request().Context(), "SalesorderDeliveryREST:CreateSalesorder")
	defer trace.Finish()

	var payload shareddomain.Salesorder
	if err := c.Bind(&payload); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	err := h.uc.Salesorder().CreateSalesorder(ctx, &payload)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "Success").JSON(c.Response())
}

func (h *RestHandler) updateSalesorder(c echo.Context) error {
	trace, ctx := tracer.StartTraceWithContext(c.Request().Context(), "SalesorderDeliveryREST:UpdateSalesorder")
	defer trace.Finish()

	var payload shareddomain.Salesorder
	if err := c.Bind(&payload); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	err := h.uc.Salesorder().UpdateSalesorder(ctx, c.Param("id"), &payload)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "Success").JSON(c.Response())
}

func (h *RestHandler) deleteSalesorder(c echo.Context) error {
	trace, ctx := tracer.StartTraceWithContext(c.Request().Context(), "SalesorderDeliveryREST:DeleteSalesorder")
	defer trace.Finish()

	if err := h.uc.Salesorder().DeleteSalesorder(ctx, c.Param("id")); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "Success").JSON(c.Response())
}