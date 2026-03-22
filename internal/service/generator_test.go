package service

import (
	"testing"
)

type mockRepo struct {
	existing map[string]bool
}

func (m *mockRepo) Check(shortLink string) (bool, error) {
	_, ok := m.existing[shortLink]
	return ok, nil
}

func (m *mockRepo) Save(originalLink, shortLink string) (string, error) {
	if m.existing == nil {
		m.existing = make(map[string]bool)
	}
	m.existing[shortLink] = true
	return shortLink, nil
}

func (m *mockRepo) Get(shortLink string) (string, error) {
	return "", nil
}

func (m *mockRepo) GetByOriginal(originalLink string) (string, error) {
	return "", nil
}

// Test generated URL len
func TestGenerateShortURL_Length(t *testing.T) {
	repo := &mockRepo{existing: make(map[string]bool)}
	gen := NewGenerator(repo)

	short, err := gen.GenerateShortURL("https://google.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(short) != 10 {
		t.Errorf("expected length 10, got %d (%s)", len(short), short)
	}
}

// Test generated URL determ
func TestGenerateShortURL_Deterministic(t *testing.T) {
	repo := &mockRepo{existing: make(map[string]bool)}
	gen := NewGenerator(repo)

	url := "https://google.com"
	short1, _ := gen.GenerateShortURL(url)
	short2, _ := gen.GenerateShortURL(url)

	if short1 != short2 {
		t.Errorf("same URL should produce same short link\n got1: %s\n got2: %s", short1, short2)
	}
}

// Test the defference between generated URLs
func TestGenerateShortURL_DifferentURLs(t *testing.T) {
	repo := &mockRepo{existing: make(map[string]bool)}
	gen := NewGenerator(repo)

	short1, _ := gen.GenerateShortURL("https://google.com")
	short2, _ := gen.GenerateShortURL("https://facebook.com")

	if short1 == short2 {
		t.Errorf("different URLs should produce different short links, but got same: %s", short1)
	}
}

// Test the permition of generated URL
func TestGenerateShortURL_AllowedChars(t *testing.T) {
	repo := &mockRepo{existing: make(map[string]bool)}
	gen := NewGenerator(repo)

	short, err := gen.GenerateShortURL("https://google.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for i, ch := range short {
		allowed := (ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '_'

		if !allowed {
			t.Errorf("position %d: character '%c' is not allowed in short URL", i, ch)
		}
	}
}

// Collision test
func TestGenerateShortURL_WithCollision(t *testing.T) {
	repo := &mockRepo{
		existing: map[string]bool{
			"abc123def0": true,
		},
	}
	gen := NewGenerator(repo)

	short, err := gen.GenerateShortURL("https://google.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if short == "" {
		t.Error("should generate some link even with collision")
	}

	if short == "abc123def0" {
		t.Error("generated link should not collide with existing one")
	}
}
