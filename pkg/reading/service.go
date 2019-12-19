package reading

import (
	"github.com/cottrellio/cottrellio_go/pkg/db"
	"github.com/cottrellio/cottrellio_go/pkg/model"
)

// Service dictates how to interface with CREATE operations.
type Service interface {
	UserList(map[string][]string) ([]*model.User, int64, error)
	UserDetail(string) (*model.User, error)
	PostList(map[string][]string) ([]*model.Post, int64, error)
	PostDetail(string) (*model.Post, error)
}

type service struct {
	db db.DB
}

// NewService creates a notifying service with the necessary dependencies.
func NewService(db db.DB) Service {
	return &service{db}
}

// BuildFiltersAndOptions builds filters and options from params.
func (s *service) BuildFiltersAndOptions(params map[string][]string) (map[string][]string, map[string]string, error) {
	filters := map[string][]string{}
	options := map[string]string{}

	// for k, vals := range params {

	// }

	return filters, options, nil
}
