package entity

import "github.com/google/uuid"

type Task struct {
	CompletedAt			*string `json:"completed_at"`
	CreatedAt       string `json:"created_at"`
	CreatedByUserID uuid.UUID `json:"created_by_user_id"`
	ID              uuid.UUID `json:"id"`
	Text            string `json:"text"`
}
