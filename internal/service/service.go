package service

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"strings"

	"github.com/Serioga111/CutterService/internal/repositorie"
)

type Generator struct {
	repo repositorie.Repositorie
}

func NewGenerator(repo repositorie.Repositorie) *Generator {
	return &Generator{repo: repo}
}

func (g *Generator) GenerateShortURL(originalUrl string) (string, error) {
	hash := sha256.Sum256([]byte(originalUrl))
	short := hashtoStr(hash[:])

	exists, err := g.repo.Check(short)
	if err != nil {
		return "", err
	}

	if !exists {
		return short, nil
	}

	for i := 1; i <= 9; i++ {
		salted := []byte(originalUrl + string(rune(i)))
		hash = sha256.Sum256(salted)
		short := hashtoStr(hash[:])
		exists, err := g.repo.Check(short)
		if err != nil {
			return "", err
		}
		if !exists {
			return short, nil
		}
	}

	return randomString(), nil
}

func hashtoStr(data []byte) string {
	result := base64.RawURLEncoding.EncodeToString(data[:])
	// Заменяем - на _ (дефис на подчеркивание)
	result = strings.ReplaceAll(result, "-", "_")
	return result[:10]
}

func randomString() string {
	const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	result := make([]byte, 10)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
