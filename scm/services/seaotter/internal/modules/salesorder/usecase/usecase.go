// Code generated by candi v1.8.8.

package usecase

import (
	"context"

	"monorepo/services/seaotter/internal/modules/salesorder/domain"
	shareddomain "monorepo/services/seaotter/pkg/shared/domain"

	"github.com/Bhinneka/candi/candishared"
)

// SalesorderUsecase abstraction
type SalesorderUsecase interface {
	GetAllSalesorder(ctx context.Context, filter *domain.FilterSalesorder) (data []shareddomain.Salesorder, meta candishared.Meta, err error)
	GetDetailSalesorder(ctx context.Context, id string) (data shareddomain.Salesorder, err error)
	CreateSalesorder(ctx context.Context, data *shareddomain.Salesorder) (err error)
	UpdateSalesorder(ctx context.Context, id string, data *shareddomain.Salesorder) (err error)
	DeleteSalesorder(ctx context.Context, id string) (err error)
}