package reading

import "github.com/cottrellio/cottrellio_go/pkg/model"

// PostList creates a user.
func (s *service) PostList(params map[string][]string) ([]*model.Post, int64, error) {
	filters, options, err := s.BuildFiltersAndOptions(params)
	if err != nil {
		return nil, -1, err
	}

	items, totalItems, err := s.db.PostList(filters, options)
	if err != nil {
		return nil, -1, err
	}

	return items, totalItems, nil
}

// PostDetail gets a user detail.
func (s *service) PostDetail(id string) (*model.Post, error) {
	post, err := s.db.PostDetail(id)
	if err != nil {
		return nil, err
	}

	return post, nil
}
