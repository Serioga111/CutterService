package repositorie

type InMemoryRepositorie struct {
}

func NewInMemoryRepositorie() *InMemoryRepositorie {
	return &InMemoryRepositorie{}
}

func (s *InMemoryRepositorie) Save(originalLink, shortLink string) error {
	return nil
}

func (s *InMemoryRepositorie) Get(shortLink string) (string, error) {
	return "", nil
}

func (s *InMemoryRepositorie) Check(shortLink string) (bool, error) {
	return false, nil
}
