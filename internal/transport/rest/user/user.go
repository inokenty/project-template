package user

import (
	"encoding/json"
	"net/http"
	"project-template/internal/iface"
	"project-template/internal/repo/model"
	"project-template/internal/transport/rest"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

type client struct {
	service iface.UserService
}

func New(s iface.UserService) iface.UserHandler {
	return &client{service: s}
}

func (c *client) AddRoutes(router *chi.Mux) {
	router.Post("/users", c.CreateUser)
	router.Get("/users/{id}", c.GetUserByID)
	router.Get("/users", c.ListUsers)
}

func (c *client) CreateUser(w http.ResponseWriter, r *http.Request) {
	var args CreateUserArgs

	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		rest.ReplyError(w, http.StatusBadRequest, errors.New("json.Decode"))

		return
	}

	id, err := c.service.CreateUser(&model.User{
		Email: args.Email,
		Phone: args.Phone,
	})
	if err != nil {
		rest.ReplyError(w, http.StatusInternalServerError, errors.Wrap(err, "service.CreateUser"))

		return
	}

	rest.ReplySuccess(w, http.StatusOK, CreateUserReply{
		ID: id,
	})
}

func (c *client) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		rest.ReplyError(w, http.StatusBadRequest, errors.Wrap(err, "strconv.ParseUint"))

		return
	}

	user, err := c.service.GetUserByID(uint(userID))
	if err != nil {
		rest.ReplyError(w, http.StatusInternalServerError, errors.Wrap(err, "service.GetUserByID"))

		return
	}

	rest.ReplySuccess(w, http.StatusOK, &GetUserByIDReply{
		User: &User{
			ID:    user.ID,
			Email: user.Email,
			Phone: user.Phone,
		},
	})
}

func (c *client) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.service.ListUsers()
	if err != nil {
		rest.ReplyError(w, http.StatusInternalServerError, errors.Wrap(err, "service.ListUsers"))

		return
	}

	items := make([]*User, 0, len(users))

	for _, user := range users {
		items = append(items, &User{
			ID:    user.ID,
			Email: user.Email,
			Phone: user.Phone,
		})
	}

	rest.ReplySuccess(w, http.StatusOK, ListUsersReply{
		Items: items,
	})
}
