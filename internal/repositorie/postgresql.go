package storage

type PostgresStorage struct {
}

func NewPostgresStorage() *PostgresStorage {
	return &PostgresStorage{}
}

func (s *PostgresStorage) Save(originalLink, shortLink string) error

func (s *PostgresStorage) Get(shortLink string) (string, error)

func (s *PostgresStorage) Check(shortLink string) (bool, error)
