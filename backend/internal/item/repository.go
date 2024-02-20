package item

import (
	"context"

	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type repository interface {
	Get(ctx context.Context, id string) (entity.Item, error)
	Create(ctx context.Context, item entity.Item) error
}

type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) Get(ctx context.Context, id string) (entity.Item, error) {
	var item entity.Item
	err := r.db.With(ctx).Select().Model(id, &item)
	return item, err
}

func (r repository) Create(ctx context.Context, item entity.Item) error {
	return r.db.With(ctx).Model(&item).Insert()
}
