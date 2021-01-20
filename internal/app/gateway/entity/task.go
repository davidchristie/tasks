package entity

import "github.com/google/uuid"

type Task struct {
	CreatedAt       string
	CreatedByUserID uuid.UUID
	ID              uuid.UUID
	Text            string
}
