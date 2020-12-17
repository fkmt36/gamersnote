package session

func (s service) Delete(key string) {
	s.c.Delete(key)
}
