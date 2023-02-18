package password

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash function hashes password using bcrypt algorithm
//
// bcrypt cost located at .env
func Hash(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

// Compare function compares a bcrypt hashed password with its possible plaintext equivalent. 
// Returns true if it's equal
func Compare(hashedPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}