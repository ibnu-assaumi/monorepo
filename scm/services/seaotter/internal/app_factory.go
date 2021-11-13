package seaotter

import (
	"monorepo/globalshared"
	"net/http"

	"github.com/Bhinneka/candi/candishared"
	restserver "github.com/Bhinneka/candi/codebase/app/rest_server"
	"github.com/Bhinneka/candi/codebase/factory"
	"github.com/Bhinneka/candi/config/env"
)

// custom from default config in candi https://github.com/Bhinneka/candi/tree/master/codebase/factory/appfactory

/*
newAppFromEnvironmentConfig constructor

Construct server/worker for running application from environment value

## Server

USE_REST=[bool]

USE_GRPC=[bool]

USE_GRAPHQL=[bool]

## Worker

USE_KAFKA_CONSUMER=[bool] # event driven handler

USE_CRON_SCHEDULER=[bool] # static scheduler

USE_REDIS_SUBSCRIBER=[bool] # dynamic scheduler

USE_TASK_QUEUE_WORKER=[bool]

USE_POSTGRES_LISTENER_WORKER=[bool]

USE_RABBITMQ_CONSUMER=[bool] # event driven handler and dynamic scheduler
*/
func newAppFromEnvironmentConfig(service factory.ServiceFactory) (apps []factory.AppServerFactory) {

	// if env.BaseEnv().UseKafkaConsumer {
	// 	apps = append(apps, setupKafkaWorker(service))
	// }
	// if env.BaseEnv().UseCronScheduler {
	// 	apps = append(apps, setupCronWorker(service))
	// }
	// if env.BaseEnv().UseTaskQueueWorker {
	// 	apps = append(apps, setupTaskQueueWorker(service))
	// }
	// if env.BaseEnv().UseRedisSubscriber {
	// 	apps = append(apps, setupRedisWorker(service))
	// }
	// if env.BaseEnv().UsePostgresListenerWorker {
	// 	apps = append(apps, setupPostgresWorker(service))
	// }
	// if env.BaseEnv().UseRabbitMQWorker {
	// 	apps = append(apps, setupRabbitMQWorker(service))
	// }

	if env.BaseEnv().UseREST {
		apps = append(apps, setupRESTServer(service))
	}
	// if env.BaseEnv().UseGRPC {
	// 	apps = append(apps, setupGRPCServer(service))
	// }
	// if !env.BaseEnv().UseREST && env.BaseEnv().UseGraphQL {
	// 	apps = append(apps, setupGraphQLServer(service))
	// }

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
