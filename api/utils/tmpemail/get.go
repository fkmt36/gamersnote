package tmpemail

func (s service) Get(uid string) *Data {
	data, found := s.c.Get(uid)
	if !found {
		return nil
	}
	asserted, ok := data.(Data)
	if !ok {
		return nil
	}
	return &asserted
}
