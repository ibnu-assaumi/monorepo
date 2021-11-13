package usecase

import (
	"context"

	"github.com/Bhinneka/candi/tracer"
	"golang.org/x/sync/errgroup"

	"monorepo/services/seaotter/internal/modules/master/domain"
	model "monorepo/services/seaotter/pkg/shared/domain"
)

func (uc *masterUsecase) GetSOPrefix(ctx context.Context, filter domain.FilterGetAllSOPrefix) (count int64, data []model.MasterSOPrefix, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MasterUsecase:GetSOPrefix")
	defer trace.Finish()

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		count, err = uc.repoSQL.MasterRepo().CountSOPrefix(egCtx, filter)
		return err
	})
	eg.Go(func() error {
		data, err = uc.repoSQL.MasterRepo().GetAllSOPrefix(egCtx, filter)
		return err
	})

	if err := eg.Wait(); err != nil {
		return 0, nil, err
	}

	return count, data, nil
}
