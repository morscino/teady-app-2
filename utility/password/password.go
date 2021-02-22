package password

import (
	"crypto/rand"
	"crypto/subtle"
	"strings"

	"golang.org/x/crypto/scrypt"
)

// password constants
const (
	SaltLen = 32
	HashLen = 64
)

type Hash struct {
	Hash []byte
	Salt []byte
}

func NewPasswordHash(password string) (*Hash, error) {
	salt := generateSalt()
	hash, err := createPasswordHash(password, salt)
	if err != nil {
		return nil, err
	}
	return &Hash{Hash: hash, Salt: salt}, nil
}

//Generate salt
func generateSalt() []byte {
	salt := make([]byte, SaltLen)

	_, _ = rand.Read(salt)
	return salt
}

func createPasswordHash(password string, salt []byte) ([]byte, error) {
	password = strings.TrimSpace(password)
	hash, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, HashLen)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// VerifyPassword checks that a password matches a stored hash and salt
func VerifyPassword(password string, hash []byte, salt []byte) bool {
	newPass, err := createPasswordHash(password, salt)
	if err != nil {
		return false
	}
	return subtle.ConstantTimeCompare(newPass, hash) == 1
}
