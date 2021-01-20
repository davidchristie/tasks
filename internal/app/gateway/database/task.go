package database

import (
	"context"
	"database/sql"

	"github.com/davidchristie/tasks/internal/app/gateway/entity"
	"github.com/google/uuid"
)

func (d *database) FindTasksCreatedByUserID(userID uuid.UUID, limit int) ([]*entity.Task, error) {
	const query = `
		SELECT id, text, created_at, created_by_user_id FROM tasks
		WHERE created_by_user_id = $1
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
		err = rows.Scan(&task.ID, &task.Text, &task.CreatedAt, &task.CreatedByUserID)
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
		INSERT INTO tasks (id, text, created_at, created_by_user_id)
		VALUES ($1, $2, $3, $4);
	`

	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, task.ID, task.Text, task.CreatedAt, task.CreatedByUserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
