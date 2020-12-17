package session

func (s service) Get(key string) *Data {

	// メモリを検索
	data, found := s.c.Get(key)
	if !found {
		return nil
	}

	// sessionDataに変換
	asserted, ok := data.(Data)
	if !ok {
		return nil
	}

	return &asserted
}
