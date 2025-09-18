package mock

import (
	"time"

	"github.com/phuocnov/golang-webserver/pkg/models"
)

var mockUsers = &models.User{
	ID:      1,
	Name:    "John Doe",
	Email:   "johndoe@gmail.com",
	Created: time.Now(),
}

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "johndoe@gmail.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "johndoe@gmail.com":
		return mockUsers.ID, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUsers, nil
	default:
		return nil, models.ErrNoRecord
	}
}
