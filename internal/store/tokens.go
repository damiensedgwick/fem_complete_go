package store

import (
	"database/sql"
	"fem_complete_go/internal/tokens"
	"fmt"
	"time"
)

type SqliteTokenStore struct {
	db *sql.DB
}

func NewSqliteTokenStore(db *sql.DB) *SqliteTokenStore {
	return &SqliteTokenStore{
		db: db,
	}
}

type TokenStore interface {
	Insert(token *tokens.Token) error
	CreateNewToken(userID int64, ttl time.Duration, scope string) (*tokens.Token, error)
	DeleteAllTokensForUser(userID int64, scope string) error
}

func (s *SqliteTokenStore) CreateNewToken(userID int64, ttl time.Duration, scope string) (*tokens.Token, error) {
	token, err := tokens.GenerateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	fmt.Println("CREATE TOKEN", token.Hash)
	err = s.Insert(token)

	return token, err
}

func (s *SqliteTokenStore) Insert(token *tokens.Token) error {
	fmt.Println("INSERT TOKEN", token.Hash)
	query := `
		INSERT INTO tokens (hash, user_id, expires_at, scope)
		VALUES (?, ?, ?, ?)
	`
	_, err := s.db.Exec(query, token.Hash, token.UserID, token.Expiry, token.Scope)
	return err
}

func (s *SqliteTokenStore) DeleteAllTokensForUser(userID int64, scope string) error {
	query := `
		DELETE FROM tokens
		WHERE user_id = ? AND scope = ?
	`
	_, err := s.db.Exec(query, userID, scope)
	return err
}
