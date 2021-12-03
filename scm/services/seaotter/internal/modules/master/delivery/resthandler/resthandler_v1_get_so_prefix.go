package resthandler

import (
	"net/http"

	"github.com/Bhinneka/candi/candishared"
	"github.com/Bhinneka/candi/tracer"
	"github.com/Bhinneka/candi/wrapper"
	"github.com/labstack/echo"

	"monorepo/globalshared"
	"monorepo/services/seaotter/internal/modules/master/payload"
)

func (r *RestHandler) getSOPrefix(c echo.Context) error {
	trace, ctx := tracer.StartTraceWithContext(c.Request().Context(), "RestHandler:getSOPrefix")
	defer trace.Finish()

	var request payload.RequestGetSOPrexif
	if err := globalshared.EchoValidateQueryParam(c, &request, "api/jsonschema/master/get_so_prefix", r.validator); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	totalData, data, err := r.uc.Master().GetSOPrefix(ctx, request)
	if err != nil {
		return wrapper.NewHTTPResponse(globalshared.GetErrorResponse(err)).JSON(c.Response())
	}

	page, limit := globalshared.EchoGetPageLimit(request.Filter, int(totalData))
	meta := candishared.NewMeta(page, limit, int(totalData))

	return wrapper.NewHTTPResponse(http.StatusOK, "success", meta, data).JSON(c.Response())
}
