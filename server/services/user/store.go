package user

import (
	"context"
	"database/sql"

	"github.com/MahatVasudev/liveStreamingProject/server/typestore"
	"github.com/redis/go-redis/v9"
)

type Store struct {
	pqdb *sql.DB
	redb *redis.Client
}

func NewStore(pqdb *sql.DB, redb *redis.Client) *Store {
	return &Store{
		pqdb,
		redb,
	}
}

func (s *Store) GetUserByID(ctx context.Context, id string) (*typestore.User, error) {
	panic("unimplemented")
}
