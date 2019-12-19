package updating

import "github.com/cottrellio/cottrellio_go/pkg/model"

// PostUpdate creates a post.
func (s *service) PostUpdate(id string, p model.Post) (*model.Post, error) {
	created, err := s.db.PostUpdate(id, p)
	if err != nil {
		return nil, err
	}

	return created, nil
}
