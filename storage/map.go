package storage

func (s *Storage) SetMap(key string, value map[string]string, expirationSec int64) {
	s.set(key, value, expirationSec)
}

func (s *Storage) GetMap(key string) (map[string]string, error) {
	val := s.get(key)

	if item, ok := val.(map[string]string); ok {
		return item, nil
	}

	return nil, newErrCustom(errWrongType)
}

func (s *Storage) GetMapItem(key, itemKey string) (string, error) {
	shard := s.getShard(key)

	shard.mutex.Lock()
	defer shard.mutex.Unlock()

	if item, ok := shard.keyValues[key]; ok {
		if isExpired(item.Expiration) {
			return "", newErrCustom(errNotExist)
		}

		if m, ok := item.Value.(map[string]string); ok {
			return m[itemKey], nil
		}
		return "", newErrCustom(errWrongType)
	}
	return "", newErrCustom(errNotExist)
}
