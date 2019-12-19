package deleting

// PostDelete creates a post.
func (s *service) PostDelete(id string) error {
	err := s.db.PostDelete(id)
	if err != nil {
		return err
	}

	return nil
}
