package auth

import "testing"

func TestHashAndCompare(t *testing.T) {
	pwd := "hovnokleslo"

	hash, err := HashPassword(pwd)
	if err != nil {
		t.Fatal("Failed to hash a password")
	}

	is_ok := CheckPasswordHash(pwd, hash)
	if !is_ok {
		t.Fatal("Hash doesn't match the password")
	}
}
