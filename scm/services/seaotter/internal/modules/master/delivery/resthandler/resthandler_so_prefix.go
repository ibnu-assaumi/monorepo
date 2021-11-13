package resthandler

import (
	"monorepo/globalshared"
	"monorepo/services/seaotter/internal/modules/master/domain"
	"net/http"

	"github.com/Bhinneka/candi/candishared"
	"github.com/Bhinneka/candi/tracer"
	"github.com/Bhinneka/candi/wrapper"
	"github.com/labstack/echo"
)

func (h *RestHandler) GetSOPrefix(c echo.Context) error {
	trace, ctx := tracer.StartTraceWithContext(c.Request().Context(), "RestHandler:GetSOPrefix")
	defer trace.Finish()

	var filter domain.FilterGetAllSOPrefix
	if err := globalshared.EchoValidateQueryParam(c, &filter, "get/master/api/v2/soprefix", h.validator); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	totalData, data, err := h.uc.Master().GetSOPrefix(ctx, filter)
	if err != nil {
		return wrapper.NewHTTPResponse(globalshared.GetErrorResponse(err)).JSON(c.Response())
	}

	page, limit := globalshared.EchoGetPageLimit(filter.Filter, int(totalData))
	meta := candishared.NewMeta(page, limit, int(totalData))
	return wrapper.NewHTTPResponse(http.StatusOK, "success", meta, data).JSON(c.Response())
}
