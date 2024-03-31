package users

import (
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// hashPassword creates a bcrypt hash of the password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPasswordHash compares a bcrypt hashed password with its possible plaintext equivalent
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWTToken(userID uint) (string, error) {
	// Set the token claims
	claims := &jwt.StandardClaims{
		Subject:   string(rune(userID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		IssuedAt:  time.Now().Unix(),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Define the secret key
	secretKey := []byte("your-secret-key")

	// Sign the token
	return token.SignedString(secretKey)
}
