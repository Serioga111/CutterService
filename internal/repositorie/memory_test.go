package repositorie

import (
	"testing"
)

// Test saving
func TestInMemoryRepositorie_Save(t *testing.T) {
	repo := NewInMemoryRepositorie()

	short, err := repo.Save("https://google.com", "abc123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if short != "abc123" {
		t.Errorf("expected abc123, got %s", short)
	}
}

// Test svaing duplicate
func TestInMemoryRepositorie_Save_DuplicateOriginal(t *testing.T) {
	repo := NewInMemoryRepositorie()

	short1, err := repo.Save("https://google.com", "abc123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	short2, err := repo.Save("https://google.com", "def456")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if short2 != short1 {
		t.Errorf("expected %s, got %s", short1, short2)
	}
}

// Test get
func TestInMemoryRepositorie_Get(t *testing.T) {
	repo := NewInMemoryRepositorie()
	repo.Save("https://google.com", "abc123")

	original, err := repo.Get("abc123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if original != "https://google.com" {
		t.Errorf("expected https://google.com, got %s", original)
	}
}

// Test Not found
func TestInMemoryRepositorie_Get_NotFound(t *testing.T) {
	repo := NewInMemoryRepositorie()

	original, err := repo.Get("notexist")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if original != "" {
		t.Errorf("expected empty string, got %s", original)
	}
}

// Test check
func TestInMemoryRepositorie_Check(t *testing.T) {
	repo := NewInMemoryRepositorie()
	repo.Save("https://google.com", "abc123")

	exists, err := repo.Check("abc123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !exists {
		t.Error("expected true, got false")
	}

	exists, err = repo.Check("notexist")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if exists {
		t.Error("expected false, got true")
	}
}
