package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

// JWTMaker is a JSON Web Token Maker
type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid size of secret key: must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey}, nil
}

// CreateToken create and sign a new token for a spesific username and duration
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedString, err := jwtToken.SignedString([]byte(maker.secretKey))
	return signedString, err
}

// VerifyToken will check if token is invalid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	}

	JWTtoken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	claims, ok := JWTtoken.Claims.(*Payload)
	if !ok || !JWTtoken.Valid {
		fmt.Println(err)
	}

	fmt.Printf("%v %v", claims.ID, claims.RegisteredClaims.Issuer)
	return claims, nil
}
