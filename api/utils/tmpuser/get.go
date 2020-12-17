package tmpuser

func (s service) Get(username string) *Data {
	data, found := s.c.Get(username)
	if !found {
		return nil
	}
	asserted, ok := data.(Data)
	if !ok {
		return nil
	}
	return &asserted
}
