package usecase

import (
	"context"

	"github.com/Bhinneka/candi/tracer"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"

	"monorepo/services/seaotter/internal/modules/master/domain"
	model "monorepo/services/seaotter/pkg/shared/domain"
)

func (uc *masterUsecase) GetSOPrefix(ctx context.Context, filter domain.FilterGetAllSOPrefix) (count int64, data []model.MasterSOPrefix, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MasterUsecase:GetSOPrefix")
	defer trace.Finish()
	decimal.NewFromInt(1).Div(decimal.NewFromInt(0))

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		count, err = uc.repoSQL.MasterRepo().CountSOPrefix(egCtx, filter)
		if err != nil {
			return err
		}
		return nil
	})
	eg.Go(func() error {
		data, err = uc.repoSQL.MasterRepo().GetAllSOPrefix(egCtx, filter)
		if err != nil {
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return 0, nil, err
	}

	return count, data, nil
}
