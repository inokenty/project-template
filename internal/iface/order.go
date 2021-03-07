package iface

import (
	"net/http"
	"project-template/internal/repo/model"

	"github.com/go-chi/chi"
)

type OrderRepo interface {
	CreateOrder(*model.Order) (uint, error)
	GetOrderByID(id uint) (*model.Order, error)
	ListOrders() ([]*model.Order, error)
}

type OrderService interface {
	CreateOrder(*model.Order) (uint, error)
	GetOrderByID(id uint) (*model.Order, error)
	ListOrders() ([]*model.Order, error)
}

type OrderHandler interface {
	AddRoutes(router *chi.Mux)
	CreateOrder(http.ResponseWriter, *http.Request)
	GetOrderByID(http.ResponseWriter, *http.Request)
	ListOrders(http.ResponseWriter, *http.Request)
}
