package domain

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"github.com/pkg/errors"
)

type TokenMaker interface {
	CreateToken(userID UserID, duration time.Duration) (Token, *Payload, error)
	VerifyToken(token Token) (*Payload, error)
}

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey SymmetricKey
}

func NewPasetoMaker(symmetricKey SymmetricKey) (TokenMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, errors.WithStack(fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize))
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: symmetricKey,
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(userID UserID, duration time.Duration) (Token, *Payload, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", nil, errors.WithStack(err)
	}

	pasetoToken, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	if err != nil {
		return "", nil, errors.WithStack(err)
	}

	token, err := NewToken(pasetoToken)
	if err != nil {
		return "", nil, errors.WithStack(err)
	}

	return token, payload, nil
}

func (maker *PasetoMaker) VerifyToken(token Token) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(string(token), maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, errors.Wrap(ErrInvalidToken, err.Error())
	}

	err = payload.Valid()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return payload, nil
}

type Payload struct {
	ID        SessionUUID `json:"id"`
	UserID    UserID      `json:"user_id"`
	IssuedAt  time.Time   `json:"issued_at"`
	ExpiresAt ExpiresAt   `json:"expired_at"`
}

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

func NewPayload(userID UserID, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	sessionID, err := NewSessionUUID(tokenID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	expiresAt, err := NewExpiresAt(time.Now().Add(duration))

	payload := &Payload{
		ID:        sessionID,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiresAt: expiresAt,
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(time.Time(payload.ExpiresAt)) {
		return errors.WithStack(ErrExpiredToken)
	}
	return nil
}

type SymmetricKey []byte

func NewSymmetricKeyFromString(v string) (SymmetricKey, error) {
	return SymmetricKey([]byte(v)), nil
}

type Token string

func NewToken(v string) (Token, error) {
	return Token(v), nil
}
