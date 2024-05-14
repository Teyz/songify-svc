package database_postgres

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/oklog/ulid/v2"

	"github.com/teyz/songify-svc/internal/database"
)

type DataPrefix string

const (
	Song DataPrefix = "song_"
	Game DataPrefix = "game_"
	User DataPrefix = "user_"
)

func (dp DataPrefix) String() string {
	return string(dp)
}

func (dp DataPrefix) IsValid(s string) bool {
	return strings.HasPrefix(s, string(dp)) && len(s) == len(string(dp))+ulid.EncodedSize
}

func GenerateDataPrefixWithULID[T DataPrefix](prefixType T) string {
	return string(prefixType) + ulid.Make().String()
}

type dbClient struct {
	connection *sqlx.DB
}

func NewClient(ctx context.Context, db *sqlx.DB) database.Database {
	return &dbClient{
		connection: db,
	}
}
