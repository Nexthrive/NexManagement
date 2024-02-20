package auth

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"github.com/qiangxue/go-rest-api/pkg/pagination"
)

// RegisterHandlers registers handlers for different HTTP requests.
func RegisterHandlers(rg *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	rg.Post("/login", login(service, logger))
	rg.Post("/auth/", res.Create)
	rg.Get("/auth/", res.query)
	rg.Get("/auth/<id>", res.get)

	rg.Use(authHandler)

  rg.Put("/auth/<id>", res.update)
  rg.Delete("/auth/<id>", res.delete)

}

type resource struct {
	service Service
	logger  log.Logger
}

// login returns a handler that handles user login request.
func login(service Service, logger log.Logger) routing.Handler {
	return func(c *routing.Context) error {
		var input LoginUser
		if err := c.Read(&input); err != nil {
			logger.With(c.Request.Context()).Errorf("invalid request: %v", err)
			return errors.BadRequest("")
		}
		token, err := service.Login(c.Request.Context(), input)
		if err != nil {
			return err
		}
		return c.Write(struct {
			Token string `json:"token"`
		}{token})
	}
}

func (r resource) Create(c *routing.Context) error {
	var input CreateUser
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
  
  user, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

  return c.WriteWithStatus(user, http.StatusOK)

}

func (r resource) get(c *routing.Context) error {
	user, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(user)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	pages := pagination.NewFromRequest(c.Request, pagination.DefaultPageSize)
	users, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = users
	return c.Write(pages)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateUser
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	user, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(user)
}

func (r resource) delete(c *routing.Context) error {
	user, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(user)
}
