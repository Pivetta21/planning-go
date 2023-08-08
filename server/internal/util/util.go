package util

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func ExtractPathParameter(r *http.Request, pattern string) (string, error) {
	param := chi.URLParam(r, pattern)
	if param == "" {
		return "", errors.New("malformed request")
	}

	return param, nil
}
