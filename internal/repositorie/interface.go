package repositorie

type Repositorie interface {
	Save(originalLink, shortLink string) error
	Get(shortLink string) (string, error)
	Check(shortLink string) (bool, error)
}
