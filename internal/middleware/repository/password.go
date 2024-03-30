package middlerepo

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/argon2"
)

func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)

	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(salt)[:length], nil
}

func HashPassword(password string, salt string) (string, error) {
	saltByte := []byte(salt)
	hashedPassword := argon2.IDKey([]byte(password), saltByte, 1, 64*1024, 4, 32)
	hexEncodedPassword := hex.EncodeToString(hashedPassword)
	return string(hexEncodedPassword), nil
}
