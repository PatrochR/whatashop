package model

import (
	"time"

	"github.com/google/uuid"
)

type UserID uuid.UUID

type User struct {
	ID           UserID     `json:"id"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"password_hash"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Roles        []Role
}

func NewUserID() UserID {
	return UserID(uuid.New())
}

func (id UserID) UUID() uuid.UUID {
	return uuid.UUID(id)
}
