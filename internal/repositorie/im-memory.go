package storage

type MemoryStorage struct {
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Save(originalLink, shortLink string) error

func (s *MemoryStorage) Get(shortLink string) (string, error)

func (s *MemoryStorage) Check(shortLink string) (bool, error)
