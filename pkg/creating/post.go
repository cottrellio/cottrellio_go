package creating

import "github.com/cottrellio/cottrellio_go/pkg/model"

// PostCreate creates a post.
func (s *service) PostCreate(p model.Post) (*model.Post, error) {
	created, err := s.db.PostCreate(p)
	if err != nil {
		return nil, err
	}

	return created, nil
}
