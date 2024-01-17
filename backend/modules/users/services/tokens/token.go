package tokens

import (
	"crypto/rand"
	"encoding/base64"
)

// CreateToken generates a random token.
//
// It does not take any parameters.
// It returns a string.
func CreateToken() string {
	bytes := make([]byte, 32)
	_, _ = rand.Read(bytes)

	return base64.URLEncoding.EncodeToString(bytes)
}
