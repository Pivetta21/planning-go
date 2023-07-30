package util

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"

	"github.com/google/uuid"
)

func GenerateIdentifier() string {
	truncatedUuid := uuid.New().String()[:8]
	return base64.RawURLEncoding.EncodeToString([]byte(truncatedUuid))
}

func GenerateCode() (string, error) {
	b := make([]byte, 4)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
