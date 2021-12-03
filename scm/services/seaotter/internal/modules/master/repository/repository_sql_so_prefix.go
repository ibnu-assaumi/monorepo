package repository

import (
	"context"
	"strings"

	"github.com/Bhinneka/candi/tracer"
	"gorm.io/gorm"

	"monorepo/globalshared"
	"monorepo/services/seaotter/internal/modules/master/payload"
	"monorepo/services/seaotter/pkg/constant"
	"monorepo/services/seaotter/pkg/shared/model"
)

func (r *masterRepoSQL) CountSOPrefix(ctx context.Context, request payload.RequestGetSOPrexif) (count int64, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MasterRepoSQL:CountSOPrefix")
	defer trace.Finish()

	request.ShowAll = true

	db := globalshared.SetSpanToGorm(ctx, r.readDB).Model(&model.MasterSOPrefix{})
	db = filterGetAllSOPrefix(db, request)
	if err := db.Count(&count).Error; err != nil {
		return 0, globalshared.NewDBError(err)
	}
	return count, nil
}

func (r *masterRepoSQL) GetAllSOPrefix(ctx context.Context, request payload.RequestGetSOPrexif) (data []model.MasterSOPrefix, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MasterRepoSQL:GetAllSOPrefix")
	defer trace.Finish()

	db := globalshared.SetSpanToGorm(ctx, r.readDB).Model(&model.MasterSOPrefix{})
	db = filterGetAllSOPrefix(db, request)
	if err := db.Find(&data).Error; err != nil {
		return nil, globalshared.NewDBError(err)
	}
	trace.Log(constant.TextDBResult, data)
	return data, nil
}

func filterGetAllSOPrefix(db *gorm.DB, request payload.RequestGetSOPrexif) *gorm.DB {
	if strings.TrimSpace(request.Search) != "" {
		db = db.Where(`("code" ILIKE '%' || ? || '%' OR "description" ILIKE '%' || ? || '%')`, request.Search, request.Search)
	}

	orderBy := request.OrderBy
	if strings.TrimSpace(orderBy) == "" {
		orderBy = `"modifiedAt"`
	}

	sort := request.Sort
	if strings.TrimSpace(sort) == "" {
		sort = "desc"
	}
	db = db.Order(orderBy + " " + sort)

	if !request.ShowAll {
		db = db.Offset(request.CalculateOffset()).Limit(request.Limit)
	}

	return db
}
