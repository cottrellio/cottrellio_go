package creating

import (
	"github.com/cottrellio/cottrellio_go/pkg/db"
	"github.com/cottrellio/cottrellio_go/pkg/model"
)

// Service dictates how to interface with CREATE operations.
type Service interface {
	UserCreate(model.User) (*model.User, error)
}

type service struct {
	db db.DB
}

// NewService creates a notifying service with the necessary dependencies.
func NewService(db db.DB) Service {
	return &service{db}
}

// UserCreate creates a user.
func (s *service) UserCreate(u model.User) (*model.User, error) {
	created, err := s.db.UserCreate(u)
	if err != nil {
		return nil, err
	}

	return created, nil
}
