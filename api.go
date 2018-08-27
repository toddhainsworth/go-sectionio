package sectionio

import (
	"fmt"
)

const version = "v1"

var apiBaseURL = fmt.Sprintf("https://aperture.section.io/api/%s", version)

// NewSession gives a new session struct that can be passed to
// API methods for authentication
func NewSession(username, password string) (Session, error) {
	return Session{
		Username: username,
		Password: password,
	}, nil
}

// TODO
// func encryptPassword(password string) ([]byte, error) {
// 	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// }
