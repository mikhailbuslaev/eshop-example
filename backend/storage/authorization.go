package storage

func (s *Storage) GetPassword(id string) (string, error) {
	var password string
	err := s.Db.Get(&password, `SELECT password FROM users WHERE id=$1 LIMIT 1`, id)
	if err != nil {
		return "", err
	}
	return password, nil
}