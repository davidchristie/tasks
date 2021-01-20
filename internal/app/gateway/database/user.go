package database

import (
	"context"
	"database/sql"

	"github.com/davidchristie/tasks/internal/app/gateway/entity"
	"github.com/google/uuid"
)

func (d *database) InsertUser(ctx context.Context, user *entity.User) error {
	const query = `
		INSERT INTO users (id, name, email, github_id, avatar_url)
		VALUES ($1, $2, $3, $4, $5);
	`

	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, user.ID, user.Name, user.Email, user.GithubID, user.AvatarURL)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (d *database) FindUserByGithubID(githubID int) (*entity.User, error) {
	const query = `
		SELECT id, name, email, github_id, avatar_url FROM users
		WHERE github_id = $1
	`

	row := d.db.QueryRow(query, githubID)
	user := entity.User{}
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.GithubID, &user.AvatarURL); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (d *database) FindUserByID(id uuid.UUID) (*entity.User, error) {
	const query = `
		SELECT id, name, email, github_id, avatar_url FROM users
		WHERE id = $1
	`

	row := d.db.QueryRow(query, id)
	user := entity.User{}
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.GithubID, &user.AvatarURL); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (d *database) HasUserWithGithubID(githubID int) (bool, error) {
	_, err := d.FindUserByGithubID(githubID)
	if err == ErrNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
