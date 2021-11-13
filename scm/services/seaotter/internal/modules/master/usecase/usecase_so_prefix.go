package usecase

import (
	"context"
	"strings"

	"github.com/Bhinneka/candi/tracer"
	"golang.org/x/sync/errgroup"

	"monorepo/services/seaotter/internal/modules/master/domain"
	"monorepo/services/seaotter/pkg/constant"
	model "monorepo/services/seaotter/pkg/shared/domain"
)

func (uc *masterUsecase) GetSOPrefix(ctx context.Context, filter domain.FilterGetAllSOPrefix) (count int64, data []model.MasterSOPrefix, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MasterUsecase:GetSOPrefix")
	defer trace.Finish()

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		count, err = uc.repoSQL.MasterRepo().CountSOPrefix(egCtx, filter)
		if err != nil && !strings.Contains(err.Error(), constant.ErrContextCancelled) {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		data, err = uc.repoSQL.MasterRepo().GetAllSOPrefix(egCtx, filter)
		if err != nil && !strings.Contains(err.Error(), constant.ErrContextCancelled) {
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return 0, nil, err
	}

	return count, data, nil
}
