package database

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/golang-migrate/migrate/v4/source/file" // file driver
	"github.com/google/uuid"
	_ "github.com/lib/pq" // postgres driver
)

type Database interface {
	FindUserByGithubID(githubID int) (*User, error)
	HasUserWithGithubID(githubID int) (bool, error)
	InsertUser(ctx context.Context, user *User) error
	Migrate(url string) error
}

type User struct {
	Email    string
	ID       uuid.UUID
	Name     string
	GithubID int
}

type database struct {
	db *sql.DB
}

var ErrNotFound = errors.New("not found")

func Connect(source string) (Database, error) {
	db, err := sql.Open("postgres", source)
	if err != nil {
		return nil, err
	}
	return &database{db}, nil
}
