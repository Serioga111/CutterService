// internal/repositorie/interface.go
package repositorie

type Repositorie interface {
	Save(originalLink, shortLink string) (string, error) // возвращает shortLink
	Get(shortLink string) (string, error)
	Check(shortLink string) (bool, error)
}
