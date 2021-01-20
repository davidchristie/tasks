package database

import (
	"context"
	"database/sql"

	"github.com/davidchristie/tasks/internal/app/gateway/entity"
	_ "github.com/golang-migrate/migrate/v4/source/file" // file driver
	"github.com/google/uuid"
	_ "github.com/lib/pq" // postgres driver
)

type Database interface {
	DeleteTask(id uuid.UUID) error
	FindTaskByID(id uuid.UUID) (*entity.Task, error)
	FindTasksCreatedByUserID(userID uuid.UUID, limit int) ([]*entity.Task, error)
	FindUserByGithubID(githubID int) (*entity.User, error)
	FindUserByID(id uuid.UUID) (*entity.User, error)
	HasUserWithGithubID(githubID int) (bool, error)
	InsertTask(ctx context.Context, task *entity.Task) error
	InsertUser(ctx context.Context, user *entity.User) error
	Migrate(url string) error
}

type database struct {
	db *sql.DB
}

func Connect(url string) (Database, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &database{db}, nil
}
