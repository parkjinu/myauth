package util

import (
	"crypto/sha256"
	"encoding/hex"
)

// CreateHash creates a hash from a string
func CreateHash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))

	md := hash.Sum(nil)

	return hex.EncodeToString(md)
}
