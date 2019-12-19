package updating

import "github.com/cottrellio/cottrellio_go/pkg/model"

// UserUpdate creates a user.
func (s *service) UserUpdate(id string, u model.User) (*model.User, error) {
	created, err := s.db.UserUpdate(id, u)
	if err != nil {
		return nil, err
	}

	return created, nil
}
