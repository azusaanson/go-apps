package db

import (
	"context"
	"time"

	"github.com/azusaanson/invest-api/domain"
	"github.com/pkg/errors"
)

type SessionQueries interface {
	CreateSession(ctx context.Context, session *domain.Session) error
}

func (s *Store) CreateSession(
	ctx context.Context,
	session *domain.Session,
) error {
	record := &Session{
		UUID:         session.UUID().ToString(),
		UserID:       uint64(session.UserID()),
		RefreshToken: string(session.RefreshToken()),
		UserAgent:    string(session.UserAgent()),
		ClientIp:     string(session.ClientIp()),
		IsBlocked:    bool(session.IsBlocked()),
		ExpiresAt:    time.Time(session.ExpiresAt()),
	}

	if err := s.conn.Create(record).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
