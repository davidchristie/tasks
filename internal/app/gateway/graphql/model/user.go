package model

import "github.com/google/uuid"

type Task struct {
	CreatedAt       string    `json:"createdAt"`
	CreatedByUserID uuid.UUID `json:"createdBy"`
	Done						bool			`json:"done"`
	ID              string    `json:"id"`
	Text            string    `json:"text"`
}
