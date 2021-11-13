package repository

import (
	"context"
	"strings"

	"github.com/Bhinneka/candi/candishared"
	"github.com/Bhinneka/candi/tracer"
	"gorm.io/gorm"

	"monorepo/globalshared"
	"monorepo/services/seaotter/internal/modules/master/domain"
	model "monorepo/services/seaotter/pkg/shared/domain"
)

func (r *masterRepoSQL) CountSOPrefix(ctx context.Context, filter domain.FilterGetAllSOPrefix) (count int64, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MasterRepoSQL:CountSOPrefix")
	defer trace.Finish()

	filter.ShowAll = true

	db := globalshared.SetSpanToGorm(ctx, r.readDB).Model(&model.MasterSOPrefix{})
	db = filterGetAllSOPrefix(db, filter)
	if err := db.Count(&count).Error; err != nil && err != gorm.ErrRecordNotFound {
		return 0, globalshared.NewErrorDB(err.Error())
	}
	return count, nil
}

func (r *masterRepoSQL) GetAllSOPrefix(ctx context.Context, filter domain.FilterGetAllSOPrefix) (data []model.MasterSOPrefix, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MasterRepoSQL:GetAllSOPrefix")
	defer trace.Finish()

	db := globalshared.SetSpanToGorm(ctx, r.readDB).Model(&model.MasterSOPrefix{})
	db = filterGetAllSOPrefix(db, filter)
	if err := db.Find(&data).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, globalshared.NewErrorDB(err.Error())
	}
	return data, nil
}

func filterGetAllSOPrefix(db *gorm.DB, filter domain.FilterGetAllSOPrefix) *gorm.DB {
	if strings.TrimSpace(filter.Search) != "" {
		db = db.Where(`("code" ILIKE '%' || ? || '%' OR "description" ILIKE '%' || ? || '%')`, filter.Search, filter.Search)
	}
	return filterPagination(db, filter.Filter)
}

func filterPagination(db *gorm.DB, filter candishared.Filter) *gorm.DB {
	orderBy := filter.OrderBy
	if strings.TrimSpace(orderBy) == "" {
		orderBy = `"modifiedAt"`
	}

	sort := filter.Sort
	if strings.TrimSpace(sort) == "" {
		sort = "desc"
	}
	db = db.Order(orderBy + " " + sort)

	if !filter.ShowAll {
		db = db.Offset(filter.CalculateOffset()).Limit(filter.Limit)
	}

	return db

}
