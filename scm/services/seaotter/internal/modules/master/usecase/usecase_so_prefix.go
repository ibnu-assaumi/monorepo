package usecase

import (
	"context"
	"fmt"

	"github.com/Bhinneka/candi/tracer"
	"golang.org/x/sync/errgroup"

	"monorepo/globalshared"
	"monorepo/services/seaotter/internal/modules/master/payload"
	"monorepo/services/seaotter/pkg/shared/model"
)

func (uc *masterUsecase) GetSOPrefix(ctx context.Context, filter payload.RequestGetSOPrexif) (count int64, data []model.MasterSOPrefix, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MasterUsecase:GetSOPrefix")
	defer func() {
		if r := recover(); r != nil {
			err = globalshared.NewErrorInternal(fmt.Sprintf("panic GetSOPrefix : %v", r))
		}
		trace.Finish()
	}()

	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		count, err = uc.repoSQL.MasterRepo().CountSOPrefix(egCtx, filter)
		return globalshared.ParseErrorContext(err)
	})
	eg.Go(func() error {
		data, err = uc.repoSQL.MasterRepo().GetAllSOPrefix(egCtx, filter)
		return globalshared.ParseErrorContext(err)
	})
	if err := eg.Wait(); err != nil {
		return 0, nil, err
	}
	return count, data, nil
}
