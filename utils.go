package main

import (
	"crypto/rand"
	"fmt"
)

// GenerateID for getting random post ID
func GenerateID() string {
	b := make([]byte, 16)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}
