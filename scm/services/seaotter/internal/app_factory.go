package seaotter

import (
	"monorepo/globalshared"
	"net/http"

	"github.com/Bhinneka/candi/candishared"
	restserver "github.com/Bhinneka/candi/codebase/app/rest_server"
	"github.com/Bhinneka/candi/codebase/factory"
	"github.com/Bhinneka/candi/config/env"
)

func newAppFromEnvironmentConfig(service factory.ServiceFactory) (apps []factory.AppServerFactory) {
	if env.BaseEnv().UseREST {
		apps = append(apps, setupRESTServer(service))
	}
	return
}

func setupRESTServer(service factory.ServiceFactory) factory.AppServerFactory {
	return restserver.NewServer(
		service,
		restserver.SetHTTPPort(env.BaseEnv().HTTPPort),
		restserver.SetRootPath(env.BaseEnv().HTTPRootPath),
		restserver.SetIncludeGraphQL(env.BaseEnv().UseGraphQL),
		restserver.SetRootHTTPHandler(http.HandlerFunc(candishared.HTTPRoot(string(service.Name()), env.BaseEnv().BuildNumber))),
		restserver.SetSharedListener(service.GetConfig().SharedListener),
		restserver.SetDebugMode(env.BaseEnv().DebugMode),
		restserver.SetJaegerMaxPacketSize(env.BaseEnv().JaegerMaxPacketSize),
		restserver.SetRootMiddlewares(globalshared.HTTPPanicMiddleware()), // add root middleware, custom from default config in candi
	)
}
