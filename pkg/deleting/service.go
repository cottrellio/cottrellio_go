package deleting

import (
	"github.com/cottrellio/cottrellio_go/pkg/db"
)

// Service dictates how to interface with CREATE operations.
type Service interface {
	UserDelete(string) error
}

type service struct {
	db db.DB
}

// NewService creates a notifying service with the necessary dependencies.
func NewService(db db.DB) Service {
	return &service{db}
}

// UserDelete creates a user.
func (s *service) UserDelete(id string) error {
	err := s.db.UserDelete(id)
	if err != nil {
		return err
	}

	return nil
}
