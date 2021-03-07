package repo

import (
	"project-template/internal/iface"
	"project-template/internal/repo/order"
	"project-template/internal/repo/user"

	"gorm.io/gorm"
)

type Client struct {
	UserRepo  iface.UserRepo
	OrderRepo iface.OrderRepo
}

func New(db *gorm.DB) *Client {
	return &Client{
		UserRepo:  user.New(db),
		OrderRepo: order.New(db),
	}
}
