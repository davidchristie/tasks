package entity

import "github.com/google/uuid"

type User struct {
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
	ID        uuid.UUID `json:"id"`
	Name      string `json:"name"`
	GithubID  int `json:"github_id"`
}
