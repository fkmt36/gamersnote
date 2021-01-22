package tmpsession

func (s service) Get(code string) *Data {
	data, found := s.c.Get(code)
	if !found {
		return nil
	}
	asserted, ok := data.(Data)
	if !ok {
		return nil
	}
	return &asserted
}
