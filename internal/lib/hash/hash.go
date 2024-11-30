package hashlib

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(hash, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func HashUserPassword(password string) (string, error) {
	const op = "storage.HashUserPassword"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("cannot hash the password: %s, %s", op, err)
	}
	return string(hashedPassword), nil
}

func GeneratePassword() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	b := make([]rune, 6)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	return string(b)
}
