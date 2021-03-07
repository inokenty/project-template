package user

import (
	"project-template/internal/iface"
	"project-template/internal/repo/model"

	"github.com/pkg/errors"
)

type client struct {
	repo iface.UserRepo
}

func New(repo iface.UserRepo) iface.UserService {
	return &client{repo: repo}
}

func (c *client) CreateUser(user *model.User) (uint, error) {
	userID, err := c.repo.CreateUser(user)
	if err != nil {
		return 0, errors.Wrap(err, "repo.CreateUser")
	}

	return userID, nil
}

func (c *client) GetUserByID(userID uint) (*model.User, error) {
	user, err := c.repo.GetUserByID(userID)
	if err != nil {
		return nil, errors.Wrap(err, "repo.GetUserByID")
	}

	return user, nil
}

func (c *client) ListUsers() ([]*model.User, error) {
	items, err := c.repo.ListUsers()
	if err != nil {
		return nil, errors.Wrap(err, "repo.ListUsers")
	}

	return items, nil
}
