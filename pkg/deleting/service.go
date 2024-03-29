package deleting

import "github.com/cottrellio/cottrellio_go/pkg/db"

// Service dictates how to interface with CREATE operations.
type Service interface {
	UserDelete(string) error
	PostDelete(string) error
}

type service struct {
	db db.DB
}

// NewService creates a notifying service with the necessary dependencies.
func NewService(db db.DB) Service {
	return &service{db}
}
