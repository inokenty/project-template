package iface

import (
	"net/http"
	"project-template/internal/repo/model"

	"github.com/go-chi/chi"
)

type UserRepo interface {
	CreateUser(*model.User) (uint, error)
	GetUserByID(id uint) (*model.User, error)
	ListUsers() ([]*model.User, error)
}

type UserService interface {
	CreateUser(*model.User) (uint, error)
	GetUserByID(id uint) (*model.User, error)
	ListUsers() ([]*model.User, error)
}

type UserHandler interface {
	AddRoutes(router *chi.Mux)
	CreateUser(http.ResponseWriter, *http.Request)
	GetUserByID(http.ResponseWriter, *http.Request)
	ListUsers(http.ResponseWriter, *http.Request)
}
