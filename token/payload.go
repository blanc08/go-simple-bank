package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	jwt.RegisteredClaims
}

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token is expired")
)

// NewPayload creates a new token payload with a spesific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:               tokenID,
		Username:         username,
		IssuedAt:         time.Now(),
		ExpiredAt:        time.Now().Add(duration),
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	return payload, nil
}

// func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
// 	return jwt.ClaimStrings{}, nil
// }

func (payload *Payload) Valid() error {
	return nil
}

// func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
// 	return time.Now(), nil
// }
