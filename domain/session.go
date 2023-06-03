package domain

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	uuid         SessionUUID
	userID       UserID
	refreshToken Token
	userAgent    UserAgent
	clientIp     ClientIp
	isBlocked    IsBlocked
	expiresAt    ExpiresAt
}

func (s *Session) UUID() SessionUUID    { return s.uuid }
func (s *Session) UserID() UserID       { return s.userID }
func (s *Session) RefreshToken() Token  { return s.refreshToken }
func (s *Session) UserAgent() UserAgent { return s.userAgent }
func (s *Session) ClientIp() ClientIp   { return s.clientIp }
func (s *Session) IsBlocked() IsBlocked { return s.isBlocked }
func (s *Session) ExpiresAt() ExpiresAt { return s.expiresAt }

func NewSession(
	sessionUUID SessionUUID,
	userID UserID,
	refreshToken Token,
	expiresAt ExpiresAt,
	userMetaData *UserMetaData,
) (*Session, error) {
	isBlocked, _ := NewIsBlocked(false)
	return &Session{
		uuid:         sessionUUID,
		userID:       userID,
		refreshToken: refreshToken,
		userAgent:    userMetaData.UserAgent(),
		clientIp:     userMetaData.ClientIp(),
		isBlocked:    isBlocked,
		expiresAt:    expiresAt,
	}, nil
}

type SessionUUID uuid.UUID

func NewSessionUUID(v uuid.UUID) (SessionUUID, error) {
	return SessionUUID(v), nil
}

func (sessionUUID SessionUUID) ToString() string {
	return uuid.UUID(sessionUUID).String()
}

type UserAgent string

func NewUserAgent(v string) (UserAgent, error) {
	return UserAgent(v), nil
}

type ClientIp string

func NewClientIp(v string) (ClientIp, error) {
	return ClientIp(v), nil
}

type IsBlocked bool

func NewIsBlocked(v bool) (IsBlocked, error) {
	return IsBlocked(v), nil
}

type ExpiresAt time.Time

func NewExpiresAt(v time.Time) (ExpiresAt, error) {
	return ExpiresAt(v), nil
}
