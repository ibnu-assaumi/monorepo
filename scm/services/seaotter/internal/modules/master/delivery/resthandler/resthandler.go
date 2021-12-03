// Code generated by candi v1.8.8.

package resthandler

import (
	"github.com/labstack/echo"

	"github.com/Bhinneka/candi/codebase/factory/dependency"
	"github.com/Bhinneka/candi/codebase/interfaces"

	"monorepo/services/seaotter/pkg/constant"
	"monorepo/services/seaotter/pkg/shared/usecase"
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
func (h *RestHandler) Mount(root *echo.Group) {
	h.mountMasterSOPrefix(root)
}

func (h *RestHandler) mountMasterSOPrefix(group *echo.Group) {
	group.GET(constant.RouteMasterSOPrefixV2, h.getSOPrefix)
}