package main

import "testing"

func TestHashAndCompare(t *testing.T) {
	pwd := "hovnokleslo"

	hash, err := hashPassword(pwd)
	if err != nil {
		t.Fatal("Failed to hash a password")
	}

	is_ok := checkPasswordHash(pwd, hash)
	if !is_ok {
		t.Fatal("Hash doesn't match the password")
	}
}
