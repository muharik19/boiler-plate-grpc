package utils

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/oklog/ulid"
)

func GeneratePK(prefix string) (*string, error) {
	// Create a source of randomness from crypto/rand.
	entropy := ulid.Monotonic(rand.Reader, 0)
	ms := ulid.Timestamp(time.Now())
	ULID, err := ulid.New(ms, entropy)
	if err != nil {
		return nil, err
	}
	id := fmt.Sprintf("%s%s", prefix, ULID.String())

	return &id, nil
}
