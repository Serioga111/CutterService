package repositorie

import (
	"sync"
)

type InMemoryRepositorie struct {
	mu      sync.RWMutex
	data    map[string]string // shortLink -> originalLink
	reverse map[string]string // originalLink -> shortLink
}

func NewInMemoryRepositorie() *InMemoryRepositorie {
	return &InMemoryRepositorie{
		data:    make(map[string]string),
		reverse: make(map[string]string),
	}
}

func (r *InMemoryRepositorie) Save(originalLink, shortLink string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if existing, ok := r.reverse[originalLink]; ok {
		return existing, nil
	}

	r.data[shortLink] = originalLink
	r.reverse[originalLink] = shortLink
	return shortLink, nil
}

func (r *InMemoryRepositorie) Get(shortLink string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	original, ok := r.data[shortLink]
	if !ok {
		return "", nil
	}
	return original, nil
}

func (r *InMemoryRepositorie) Check(shortLink string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.data[shortLink]
	return ok, nil
}
