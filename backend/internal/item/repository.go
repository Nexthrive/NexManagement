package item

import (
	"context"

	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type Repository interface {
	Get(ctx context.Context, id string) (entity.Item, error)
	Count(ctx context.Context) (int, error)

	Create(ctx context.Context, item entity.Item) error
	Delete(ctx context.Context, id string) error
	Query(ctx context.Context, offset, limit int) ([]entity.Item, error)
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

func (r repository) Delete(ctx context.Context, id string) error {
	item, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&item).Delete()
}

func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Item, error) {
	var Items []entity.Item
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&Items)
	return Items, err
}

func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("item").Row(&count)
	return count, err
}
