package handler

import (
	"project-template/internal/iface"
	"project-template/internal/service"
	"project-template/internal/transport/rest/order"
	"project-template/internal/transport/rest/user"

	"github.com/go-chi/chi"
)

type Client struct {
	UserHandler  iface.UserHandler
	OrderHandler iface.OrderHandler
}

func New(c *service.Client) *Client {
	return &Client{
		UserHandler:  user.New(c.UserService),
		OrderHandler: order.New(c.OrderService),
	}
}

func (c *Client) AddRoutes(router *chi.Mux) {
	c.UserHandler.AddRoutes(router)
	c.OrderHandler.AddRoutes(router)
}
