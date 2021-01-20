package entity

import "github.com/google/uuid"

type User struct {
	AvatarURL string
	Email     string
	ID        uuid.UUID
	Name      string
	GithubID  int
}
