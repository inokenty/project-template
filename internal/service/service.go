package service

import (
	"project-template/internal/iface"
	"project-template/internal/repo"
	"project-template/internal/service/order"
	"project-template/internal/service/user"
)

type Client struct {
	UserService  iface.UserService
	OrderService iface.OrderService
}

func New(c *repo.Client) *Client {
	return &Client{
		UserService:  user.New(c.UserRepo),
		OrderService: order.New(c.OrderRepo),
	}
}
