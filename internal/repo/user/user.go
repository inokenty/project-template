package user

import (
	"project-template/internal/iface"
	"project-template/internal/repo/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type client struct {
	db *gorm.DB
}

func New(db *gorm.DB) iface.UserRepo {
	return &client{db: db}
}

func (c *client) CreateUser(user *model.User) (uint, error) {
	if err := c.db.Create(user).Error; err != nil {
		return 0, errors.Wrap(err, "db.Create")
	}

	return user.ID, nil
}

func (c *client) GetUserByID(userID uint) (*model.User, error) {
	var user model.User

	if err := c.db.First(&user, userID).Error; err != nil {
		return nil, errors.Wrap(err, "db.First")
	}

	return &user, nil
}

func (c *client) ListUsers() ([]*model.User, error) {
	var users []*model.User

	if err := c.db.Find(&users).Error; err != nil {
		return nil, errors.Wrap(err, "db.Find")
	}

	return users, nil
}
