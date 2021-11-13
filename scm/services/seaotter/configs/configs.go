// Code generated by candi v1.8.8.

package configs

import (
	"context"

	"monorepo/sdk"
	"monorepo/services/seaotter/pkg/shared"
	"monorepo/services/seaotter/pkg/shared/repository"
	"monorepo/services/seaotter/pkg/shared/usecase"

	"github.com/Bhinneka/candi/broker"
	"github.com/Bhinneka/candi/candihelper"
	"github.com/Bhinneka/candi/codebase/factory/dependency"
	"github.com/Bhinneka/candi/codebase/interfaces"
	"github.com/Bhinneka/candi/config"
	"github.com/Bhinneka/candi/config/database"
	"github.com/Bhinneka/candi/logger"
	"github.com/Bhinneka/candi/middleware"
	"github.com/Bhinneka/candi/tracer"
	"github.com/Bhinneka/candi/validator"
)

// LoadServiceConfigs load selected dependency configuration in this service
func LoadServiceConfigs(baseCfg *config.Config) (deps dependency.Dependency) {
	logger.InitZap()

	var sharedEnv shared.Environment
	candihelper.MustParseEnv(&sharedEnv)
	shared.SetEnv(sharedEnv)

	tracer.InitOpenTracing(baseCfg.ServiceName)

	baseCfg.LoadFunc(func(ctx context.Context) []interfaces.Closer {
		brokerDeps := broker.InitBrokers(
			broker.NewKafkaBroker(),
			// broker.NewRabbitMQBroker(),
		)
		redisDeps := database.InitRedis()
		sqlDeps := database.InitSQLDatabase()
		mongoDeps := database.InitMongoDB(ctx)

		sdk.SetGlobalSDK(
		// init service client sdk
		)

		// inject all service dependencies
		// See all option in dependency package
		deps = dependency.InitDependency(
			dependency.SetMiddleware(middleware.NewMiddleware(
				&shared.DefaultTokenValidator{},
				&shared.DefaultACLPermissionChecker{}),
			),
			dependency.SetValidator(validator.NewValidator()),
			dependency.SetBrokers(brokerDeps.GetBrokers()),
			dependency.SetRedisPool(redisDeps),
			dependency.SetSQLDatabase(sqlDeps),
			dependency.SetMongoDatabase(mongoDeps),
			// ... add more dependencies
		)
		return []interfaces.Closer{ // throw back to base config for close connection when application shutdown
			brokerDeps,
			redisDeps,
			sqlDeps,
			mongoDeps,
		}
	})

	repository.SetSharedRepository(deps)
	usecase.SetSharedUsecase(deps)

	return deps
}
