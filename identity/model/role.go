package model

import (
	"github.com/google/uuid"
)

type RoleID uuid.UUID

type Role struct {
	ID   RoleID
	Name string
}

func NewRoleID() RoleID{
	return RoleID(uuid.New())
}

func(id RoleID) UUID() uuid.UUID {
	return uuid.UUID(id)
}
