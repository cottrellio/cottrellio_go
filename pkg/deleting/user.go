package deleting

// UserDelete creates a user.
func (s *service) UserDelete(id string) error {
	err := s.db.UserDelete(id)
	if err != nil {
		return err
	}

	return nil
}
