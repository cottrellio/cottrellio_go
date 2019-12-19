package creating

import "github.com/cottrellio/cottrellio_go/pkg/model"

// UserCreate creates a user.
func (s *service) UserCreate(u model.User) (*model.User, error) {
	created, err := s.db.UserCreate(u)
	if err != nil {
		return nil, err
	}

	return created, nil
}
