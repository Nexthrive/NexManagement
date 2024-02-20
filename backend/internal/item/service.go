package item

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type Service interface {
	Create(ctx context.Context, Input CreateItemReq) (Item, error)
	Get(ctx context.Context, id string) (Item, error)
	Count(ctx context.Context) (int, error)
	Update(ctx context.Context, id string, input UpdateItemReq) (Item, error)
	Query(ctx context.Context, offset, limit int) ([]Item, error)
	Delete(ctx context.Context, id string) (string, error)
}

type Item struct {
	entity.Item
}

type CreateItemReq struct {
	Name      string `json:"name"`
	Deskripsi string `json:"Deskripsi"`
	Stok      int    `json:"stok"`
	Harga     int    `json:"harga"`
}

func (m CreateItemReq) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Deskripsi, validation.Required, validation.Length(0, 500)),
		validation.Field(&m.Stok, validation.Required, validation.Min(0)),
		validation.Field(&m.Harga, validation.Required, validation.Min(0)),
	)
}

type UpdateItemReq struct {
	Name      string `json:"name"`
	Deskripsi string `json:"Deskripsi"`
	Stok      int    `json:"stok"`
	Harga     int    `json:"harga"`
}

func (m UpdateItemReq) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Deskripsi, validation.Required, validation.Length(0, 500)),
		validation.Field(&m.Stok, validation.Required, validation.Min(0)),
		validation.Field(&m.Harga, validation.Required, validation.Min(0)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return &service{repo, logger}
}

func (s *service) Get(ctx context.Context, id string) (Item, error) {
	item, err := s.repo.Get(ctx, id)
	if err != nil {
		return Item{}, err
	}
	return Item{item}, nil
}

func (s *service) Create(ctx context.Context, req CreateItemReq) (Item, error) {
	if err := req.Validate(); err != nil {
		return Item{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.Item{
		ID:        id,
		Name:      req.Name,
		Deskripsi: req.Deskripsi,
		Stok:      req.Stok,
		Harga:     req.Harga,
		CreatedAt: now,
	})
	if err != nil {
		return Item{}, err
	}
	return s.Get(ctx, id)

}

func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

func (s service) Query(ctx context.Context, offset, limit int) ([]Item, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Item{}
	for _, item := range items {
		result = append(result, Item{item})
	}
	return result, nil
}

func (s service) Update(ctx context.Context, id string, req UpdateItemReq) (Item, error) {
	if err := req.Validate(); err != nil {
		return Item{}, err
	}

	item, err := s.Get(ctx, id)
	if err != nil {
		return item, err
	}
	item.Name = req.Name
	item.Deskripsi = req.Deskripsi
	item.Stok = req.Stok
	item.Harga = req.Harga

	if err := s.repo.Update(ctx, item.Item); err != nil {
		return item, err
	}
	return item, nil
}

func (s service) Delete(ctx context.Context, id string) (string, error) {
	_, err := s.Get(ctx, id)
	if err != nil {
		return "", err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return "", err
	}
	return "Delete successful", nil
}
