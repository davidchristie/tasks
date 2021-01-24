package database

import (
	"context"
	"database/sql"

	"github.com/davidchristie/tasks/internal/app/gateway/entity"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (d *database) DeleteAllTasks() error {
	const query = `
		TRUNCATE tasks CASCADE
	`

	_, err := d.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "delete all tasks failed")
	}

	return nil
}

func (d *database) DeleteTask(id uuid.UUID) error {
	const query = `
		DELETE FROM tasks
		WHERE id = $1
	`

	res, err := d.db.Exec(query, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return ErrNotFound
	}

	return nil
}

func (d *database) FindTaskByID(id uuid.UUID) (*entity.Task, error) {
	const query = `
		SELECT id, text, completed_at, created_at, created_by_user_id FROM tasks
		WHERE id = $1
	`

	row := d.db.QueryRow(query, id)
	task := entity.Task{}
	if err := row.Scan(&task.ID, &task.Text, &task.CompletedAt, &task.CreatedAt, &task.CreatedByUserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &task, nil
}

func (d *database) FindTasksCreatedByUserID(userID uuid.UUID, limit int) ([]*entity.Task, error) {
	const query = `
		SELECT id, text, completed_at, created_at, created_by_user_id FROM tasks
		WHERE created_by_user_id = $1
		ORDER BY created_at ASC
		LIMIT $2
	`

	rows, err := d.db.Query(query, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*entity.Task{}

	for rows.Next() {
		task := entity.Task{}
		err = rows.Scan(&task.ID, &task.Text, &task.CompletedAt, &task.CreatedAt, &task.CreatedByUserID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (d *database) InsertTask(ctx context.Context, task *entity.Task) error {
	const query = `
		INSERT INTO tasks (id, text, completed_at, created_at, created_by_user_id)
		VALUES ($1, $2, $3, $4, $5);
	`

	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, task.ID, task.Text, task.CompletedAt, task.CreatedAt, task.CreatedByUserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (d *database) UpdateTask(task *entity.Task) error {
	const query = `
		UPDATE tasks
		SET text = $2, completed_at = $3, created_at = $4, created_by_user_id = $5
		WHERE id = $1;
	`

	res, err := d.db.Exec(query, task.ID, task.Text, task.CompletedAt, task.CreatedAt, task.CreatedByUserID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return ErrNotFound
	}

	return nil
}
