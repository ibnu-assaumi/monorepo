// Code generated by candi v1.8.8.

package repository

import (
	"context"
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Bhinneka/candi/candishared"
	"github.com/Bhinneka/candi/tracer"

	// @candi:repositoryImport
	"monorepo/globalshared"
	masterrepo "monorepo/services/seaotter/internal/modules/master/repository"
	salesorderrepo "monorepo/services/seaotter/internal/modules/salesorder/repository"
)

type (
	// RepoSQL abstraction
	RepoSQL interface {
		WithTransaction(ctx context.Context, txFunc func(ctx context.Context) error) (err error)

		// @candi:repositoryMethod
		MasterRepo() masterrepo.MasterRepoSQL
		SalesorderRepo() salesorderrepo.SalesorderRepository
	}

	repoSQLImpl struct {
		readDB, writeDB *gorm.DB

		// register all repository from modules
		// @candi:repositoryField
		masterRepo     masterrepo.MasterRepoSQL
		salesorderRepo salesorderrepo.SalesorderRepository
	}
)

var (
	globalRepoSQL RepoSQL
)

// setSharedRepoSQL set the global singleton "RepoSQL" implementation
func setSharedRepoSQL(readDB, writeDB *sql.DB) {
	gormRead, err := gorm.Open(postgres.New(postgres.Config{
		Conn: readDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	gormWrite, err := gorm.Open(postgres.New(postgres.Config{
		Conn: writeDB,
	}), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic(err)
	}

	globalshared.AddGormCallbacks(gormRead)
	globalshared.AddGormCallbacks(gormWrite)

	globalRepoSQL = NewRepositorySQL(gormRead, gormWrite)
}

// GetSharedRepoSQL returns the global singleton "RepoSQL" implementation
func GetSharedRepoSQL() RepoSQL {
	return globalRepoSQL
}

// NewRepositorySQL constructor
func NewRepositorySQL(readDB, writeDB *gorm.DB) RepoSQL {

	return &repoSQLImpl{
		readDB: readDB, writeDB: writeDB,

		// @candi:repositoryConstructor
		masterRepo:     masterrepo.NewMasterRepoSQL(readDB, writeDB),
		salesorderRepo: salesorderrepo.NewSalesorderRepoSQL(readDB, writeDB),
	}
}

// WithTransaction run transaction for each repository with context, include handle canceled or timeout context
func (r *repoSQLImpl) WithTransaction(ctx context.Context, txFunc func(ctx context.Context) error) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "RepoSQL:Transaction")
	defer trace.Finish()

	tx, ok := candishared.GetValueFromContext(ctx, candishared.ContextKeySQLTransaction).(*gorm.DB)
	if !ok {
		tx = r.writeDB.Begin()
		err = tx.Error
		if err != nil {
			return err
		}

		defer func() {
			if err != nil {
				tx.Rollback()
				trace.SetError(err)
			} else {
				tx.Commit()
			}
		}()

		ctx = candishared.SetToContext(ctx, candishared.ContextKeySQLTransaction, tx)
	}

	errChan := make(chan error)
	go func(ctx context.Context) {
		defer func() {
			if r := recover(); r != nil {
				errChan <- fmt.Errorf("panic: %v", r)
			}
			close(errChan)
		}()

		if err := txFunc(ctx); err != nil {
			errChan <- err
		}
	}(ctx)

	select {
	case <-ctx.Done():
		return fmt.Errorf("canceled or timeout: %v", ctx.Err())
	case e := <-errChan:
		return e
	}
}

// @candi:repositoryImplementation
func (r *repoSQLImpl) MasterRepo() masterrepo.MasterRepoSQL {
	return r.masterRepo
}

func (r *repoSQLImpl) SalesorderRepo() salesorderrepo.SalesorderRepository {
	return r.salesorderRepo
}
