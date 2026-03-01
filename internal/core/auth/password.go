package auth

import (
	"crypto/sha256"
	"encoding/hex"

	er "Board_of_issuses/internal/core"
)

const (
	MinPasswordLength int = 6
	MaxPasswordLength int = 30
)

func Hash(password string) (string, error) {

	if len(password) <= MinPasswordLength {
		return "", er.TooShortPassword()
	}

	if len(password) > MaxPasswordLength {
		return "", er.TooLongPassword()
	}

	hashPassword := sha256.Sum256([]byte(password))

	return hex.EncodeToString(hashPassword[:]), nil
}

func Compare(password, truePassword string) bool {
	hashPassword, err := Hash(password)
	if err != nil {
		return false
	}

	return hashPassword == truePassword
}
