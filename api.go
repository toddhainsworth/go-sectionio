package sectionio

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const version = "v1"

var apiBaseURL = fmt.Sprintf("https://aperture.section.io/api/%s", version)

// NewSession gives a new session struct that can be passed to
// API methods for authentication
func NewSession(username, password string) (Session, error) {
	pw, err := encryptPassword(password)

	if err != nil {
		return Session{}, err
	}

	return Session{
		Username: username,
		Password: pw,
	}, nil
}

func encryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
