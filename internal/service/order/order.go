package order

import (
	"project-template/internal/iface"
	"project-template/internal/repo/model"

	"github.com/pkg/errors"
)

type client struct {
	repo iface.OrderRepo
}

func New(repo iface.OrderRepo) iface.OrderService {
	return &client{repo: repo}
}

func (c *client) CreateOrder(order *model.Order) (uint, error) {
	orderID, err := c.repo.CreateOrder(order)
	if err != nil {
		return 0, errors.Wrap(err, "repo.CreateOrder")
	}

	return orderID, nil
}

func (c *client) GetOrderByID(orderID uint) (*model.Order, error) {
	order, err := c.repo.GetOrderByID(orderID)
	if err != nil {
		return nil, errors.Wrap(err, "repo.GetOrderByID")
	}

	return order, nil
}

func (c *client) ListOrders() ([]*model.Order, error) {
	items, err := c.repo.ListOrders()
	if err != nil {
		return nil, errors.Wrap(err, "repo.ListOrders")
	}

	return items, nil
}
