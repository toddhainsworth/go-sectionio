package sectionio

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestNewSession(t *testing.T) {
	password := []byte("password")
	session, err := NewSession("username", string(password))

	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword(session.Password, password); err != nil {
		t.Errorf("Stored password does not compare positively with the hash.")
	}
}
