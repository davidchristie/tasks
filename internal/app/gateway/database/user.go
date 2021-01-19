package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

func (d *database) InsertUser(ctx context.Context, user *User) error {
	const query = `
		INSERT INTO users (id, name, email, github_id)
		VALUES ($1, $2, $3, $4);
	`

	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, user.ID, user.Name, user.Email, user.GithubID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (d *database) FindUserByGithubID(githubID int) (*User, error) {
	const query = `
		SELECT id, name, email, github_id FROM users
		WHERE github_id = $1
	`

	row := d.db.QueryRow(query, githubID)
	var rowID uuid.UUID
	var rowName string
	var rowEmail string
	var rowGithubID int
	if err := row.Scan(&rowID, &rowName, &rowEmail, &rowGithubID); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &User{
		Email:    rowEmail,
		GithubID: rowGithubID,
		ID:       rowID,
		Name:     rowName,
	}, nil
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
