package reading

import (
	"github.com/cottrellio/cottrellio_go/pkg/db"
	"github.com/cottrellio/cottrellio_go/pkg/model"
)

// Service dictates how to interface with CREATE operations.
type Service interface {
	UserList(filters map[string][]string) ([]*model.User, int64, error)
	UserDetail(id string) (*model.User, error)
}

type service struct {
	db db.DB
}

// NewService creates a notifying service with the necessary dependencies.
func NewService(db db.DB) Service {
	return &service{db}
}

// CreateUser creates a user.
func (s *service) UserList(params map[string][]string) ([]*model.User, int64, error) {
	filters, options, err := s.BuildFiltersAndOptions(params)
	if err != nil {
		return nil, -1, err
	}

	items, totalItems, err := s.db.UserList(filters, options)
	if err != nil {
		return nil, -1, err
	}

	return items, totalItems, nil
}

// UserDetail gets a user detail.
func (s *service) UserDetail(id string) (*model.User, error) {
	user, err := s.db.UserDetail(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// BuildFiltersAndOptions builds filters and options from params.
func (s *service) BuildFiltersAndOptions(params map[string][]string) (map[string][]string, map[string]string, error) {
	filters := map[string][]string{}
	options := map[string]string{}

	// for k, vals := range params {

	// }

	return filters, options, nil
}
