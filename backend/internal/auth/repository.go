package auth

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	Get(ctx context.Context, id string) (entity.User, error)
	Create(ctx context.Context, user entity.User) error
	GetByUsernameAndPassword(ctx context.Context, Username string, passphrase string) (entity.User, error)
	Query(ctx context.Context, offset, limit int) ([]entity.User, error)
	Update(ctx context.Context, users entity.User) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

func (r repository) Create(ctx context.Context, user entity.User) error {
	return r.db.With(ctx).Model(&user).Insert()
}

func (r repository) Get(ctx context.Context, id string) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().Model(id, &user)
	return user, err
}

func (r repository) GetByUsernameAndPassword(ctx context.Context, Username string, passphrase string) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().From("user").Where(dbx.HashExp{"username": Username}).One(&user)
	if err != nil {
		return user, err
	}

	// Compare stored hash with the hash of the provided passphrase
	if err := bcrypt.CompareHashAndPassword([]byte(user.Passphrase), []byte(passphrase)); err != nil {
		return user, err
	}

	return user, nil
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.User, error) {
	var users []entity.User
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&users)
	return users, err
}

func (r repository) Update(ctx context.Context, users entity.User) error {
	return r.db.With(ctx).Model(&users).Update()
}

// Delete deletes an album with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	user, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&user).Delete()
}

func NewRepo(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}
