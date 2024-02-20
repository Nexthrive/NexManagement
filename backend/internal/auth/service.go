package auth

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(ctx context.Context, input LoginUser) (string, error)
	Create(ctx context.Context, req CreateUser) (User, error)	
  Update(ctx context.Context, id string, req UpdateUser) (User, error)
	Delete(ctx context.Context, id string) (User, error)
	Get(ctx context.Context, id string) (User, error)
	Query(ctx context.Context, offset, limit int) ([]User, error)
}

type User struct {
	entity.User
}

type service struct {
	signingKey      string
	tokenExpiration int
	logger          log.Logger
	repo            Repository
}

type CreateUser struct {
	Username   string `json:"username"`
	Passphrase string `json:"passphrase"`
	Email      string `json:"email"`
	No_telp    string `json:"no_telp"`
  Role       string `json:"role"`
}

type UpdateUser struct {
	Username   string `json:"username"`
	Passphrase string `json:"passphrase"`
	Email      string `json:"email"`
	No_telp    string `json:"no_telp"`
  Role       string `json:"role"`
}

type LoginUser struct {
	Username   string `json:"username"`
	Passphrase string `json:"passphrase"`
}

func (cu CreateUser) Validate() error {
	return validation.ValidateStruct(&cu,
		validation.Field(&cu.Username, validation.Required),
		validation.Field(&cu.Passphrase, validation.Required, validation.Length(8, 128)),
		validation.Field(&cu.Email, validation.Required),
	)
}
func (uu UpdateUser) Validate() error {
	return validation.ValidateStruct(&uu,
		validation.Field(&uu.Username, validation.Required),
		validation.Field(&uu.Passphrase, validation.Required, validation.Length(8, 128)),
		validation.Field(&uu.Email, validation.Required),
	)
}

func NewService(signingKey string, tokenExpiration int, logger log.Logger, repo Repository) Service {
	return &service{signingKey, tokenExpiration, logger, repo}
}

func (s service) Create(ctx context.Context, req CreateUser) (User, error) {
    if err := req.Validate(); err != nil {
        return User{}, err
    }
    id := entity.GenerateID()

    // Hashing passphrase
    hashedPassphrase, err := bcrypt.GenerateFromPassword([]byte(req.Passphrase), bcrypt.DefaultCost)
    if err != nil {
        return User{}, err
    }

    err = s.repo.Create(ctx, entity.User{
        ID:         id,
        Username:   req.Username,
        Passphrase: string(hashedPassphrase), // Simpan passphrase yang telah di-hash
        Email:      req.Email,
        No_telp:    req.No_telp,
        Role:       req.Role,
    })
    if err != nil {
        return User{}, err
    }
    return s.Get(ctx, id)
}

func (s *service) Login(ctx context.Context, loginReq LoginUser) (string, error) {
	user, err := s.authenticate(ctx, loginReq)
	if err != nil {
		s.logger.Error("Authentication failed", "error", err)
		return "", errors.Unauthorized("authentication failed")
	}

	return s.generateJWT(user)
}

func (s service) Update(ctx context.Context, id string, req UpdateUser) (User, error) {
	if err := req.Validate(); err != nil {
		return User{}, err
	}

	user, err := s.Get(ctx, id)
	if err != nil {
		return user, err
	}

	user.Username = req.Username
	user.Email = req.Email
	user.No_telp = req.No_telp
  user.Role = req.Role

	if req.Passphrase != "" {
		hashedPassphrase, err := bcrypt.GenerateFromPassword([]byte(req.Passphrase), bcrypt.DefaultCost)
		if err != nil {
			return User{}, err
		}
		user.Passphrase = string(hashedPassphrase)
	}

	if err := s.repo.Update(ctx, user.User); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s service) Delete(ctx context.Context, id string) (User, error) {
	user, err := s.Get(ctx, id)
	if err != nil {
		return User{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s service) Get(ctx context.Context, id string) (User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return User{}, err
	}
	return User{user}, nil
}

func (s service) Query(ctx context.Context, offset, limit int) ([]User, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []User{}
	for _, item := range items {
		result = append(result, User{item})
	}
	return result, nil
}

func (s *service) authenticate(ctx context.Context, loginReq LoginUser) (User, error) {
	user, err := s.repo.GetByUsernameAndPassword(ctx, loginReq.Username, loginReq.Passphrase)
	if err != nil {
		// If authentication fails, return unauthorized error
		s.logger.Error("Authentication failed", "error", err)
		return User{}, errors.Unauthorized("authentication failed")
	}

	return User{User: user}, nil
}

func (s service) generateJWT(user User) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"name": user.Username,
		"exp":  time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
	}).SignedString([]byte(s.signingKey))
}
