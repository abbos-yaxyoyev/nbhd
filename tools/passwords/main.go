package passwords

import (
	"golang.org/x/crypto/bcrypt"
)

func Encode(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 6)
	return string(hash)
}

func Validate(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
