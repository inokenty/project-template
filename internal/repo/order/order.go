package order

import (
	"project-template/internal/iface"
	"project-template/internal/repo/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type client struct {
	db *gorm.DB
}

func New(db *gorm.DB) iface.OrderRepo {
	return &client{db: db}
}

func (c *client) CreateOrder(order *model.Order) (uint, error) {
	if err := c.db.Create(order).Error; err != nil {
		return 0, errors.Wrap(err, "db.Create")
	}

	return order.ID, nil
}

func (c *client) GetOrderByID(id uint) (*model.Order, error) {
	var item model.Order

	if err := c.db.First(&item, id).Error; err != nil {
		return nil, errors.Wrap(err, "db.First")
	}

	return &item, nil
}

func (c *client) ListOrders() ([]*model.Order, error) {
	var items []*model.Order

	if err := c.db.Find(&items).Error; err != nil {
		return nil, errors.Wrap(err, "db.Find")
	}

	return items, nil
}
