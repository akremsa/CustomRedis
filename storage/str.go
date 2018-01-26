package storage

// SetStr sets key to hold the string value
func (s *Storage) SetStr(key, value string, expirationSec int64) {
	s.set(key, value, expirationSec)
}

// SetStrNX sets key to hold string value if key does not exist
func (s *Storage) SetStrNX(key, value string, expirationSec int64) error {
	return s.setNX(key, value, expirationSec)
}

// GetStr gets string value of the key
func (s *Storage) GetStr(key string) (string, error) {
	val := s.get(key)

	if str, ok := val.(string); ok {
		return str, nil
	}
	return "", newErrCustom(errWrongType)
}
