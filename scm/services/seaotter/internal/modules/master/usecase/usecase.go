// Code generated by candi v1.8.8.

package usecase

import (
	"context"
	"monorepo/services/seaotter/internal/modules/master/payload"
	"monorepo/services/seaotter/pkg/shared/model"
	"monorepo/services/seaotter/pkg/shared/repository"
	"monorepo/services/seaotter/pkg/shared/usecase/common"

	"github.com/Bhinneka/candi/codebase/factory/dependency"
	"github.com/Bhinneka/candi/codebase/factory/types"
	"github.com/Bhinneka/candi/codebase/interfaces"
)

type masterUsecase struct {
	sharedUsecase common.Usecase
	cache         interfaces.Cache
	repoSQL       repository.RepoSQL
	// repoMongo     repository.RepoMongo
	kafkaPub interfaces.Publisher
	// rabbitmqPub   interfaces.Publisher
}

// MasterUsecase abstraction
type MasterUsecase interface {
	GetSOPrefix(ctx context.Context, filter payload.RequestGetSOPrexif) (int64, []model.MasterSOPrefix, error)
}

// NewMasterUsecase usecase impl constructor
func NewMasterUsecase(deps dependency.Dependency) (MasterUsecase, func(sharedUsecase common.Usecase)) {
	uc := &masterUsecase{
		cache:   deps.GetRedisPool().Cache(),
		repoSQL: repository.GetSharedRepoSQL(),
		// repoMongo: repository.GetSharedRepoMongo(),
		kafkaPub: deps.GetBroker(types.Kafka).GetPublisher(),
		// rabbitmqPub: deps.GetBroker(types.RabbitMQ).GetPublisher(),
	}
	return uc, func(sharedUsecase common.Usecase) {
		uc.sharedUsecase = sharedUsecase
	}
}
