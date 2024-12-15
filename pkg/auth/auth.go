package auth

import (
    "crypto/rand"
    "encoding/base64"
    "errors"
    "strings"
    "time"

    "golang.org/x/crypto/bcrypt"
)

// User represents a user in the system
type User struct {
    ID             int
    Username       string
    PasswordHash   string
    Email          string
    LastLogin      time.Time
    CreatedAt      time.Time
}

// HashPassword generates a bcrypt hash of the password
func HashPassword(password string) (string, error) {
    // Cost of 12 provides a good balance between security and performance
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    return string(bytes), err
}

// CheckPasswordHash compares a plain text password with its hash
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// GenerateToken creates a secure random token
func GenerateToken() (string, error) {
    b := make([]byte, 32)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(b), nil
}

// ValidateEmail performs basic email validation
func ValidateEmail(email string) error {
    // Basic email validation
    if len(email) < 3 || !strings.Contains(email, "@") {
        return errors.New("invalid email format")
    }
    return nil
}

// AuthenticateUser validates user credentials
func AuthenticateUser(username, password string, userStore UserStore) (*User, error) {
    user, err := userStore.FindByUsername(username)
    if err != nil {
        return nil, errors.New("user not found")
    }

    if !CheckPasswordHash(password, user.PasswordHash) {
        return nil, errors.New("invalid credentials")
    }

    return user, nil
}

// UserStore defines the interface for user persistence
type UserStore interface {
    FindByUsername(username string) (*User, error)
    Create(user *User) error
    Update(user *User) error
}
