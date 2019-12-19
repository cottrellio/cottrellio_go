package reading

import "github.com/cottrellio/cottrellio_go/pkg/model"

// UserList creates a user.
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
