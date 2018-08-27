package sectionio

import (
	"testing"
)

func TestNewSession(t *testing.T) {
	password := "password"
	_, err := NewSession("username", password)

	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}
}
