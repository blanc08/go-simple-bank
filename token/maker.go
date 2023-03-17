package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken create and sign a new token for a spesific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken will check if token is invalid or not
	VerifyToken(token string) (*Payload, error)
}