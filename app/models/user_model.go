package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe User object.
type User struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Email        string    `db:"email" json:"email" validate:"required,email,lte=255"`
	FullName     string    `db:"full_name" json:"full_name" validate:"required,lte=25"`
	PasswordHash string    `db:"password_hash" json:"password_hash,omitempty" validate:"required,lte=255"`
	UserStatus   int       `db:"user_status" json:"user_status" validate:"required,len=1"`
}

// Role struct with UUID
type Role struct {
	ID          uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Name        string    `db:"name" json:"name" validate:"required,lte=255"`
	Description string    `db:"description" json:"description" validate:"required,lte=255"`
}

// Permission struct with UUID
type Permission struct {
	ID          uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Name        string    `db:"name" json:"name" validate:"required,lte=255"`
	Description string    `db:"description" json:"description" validate:"required,lte=255"`
}

type RolePermissions struct {
	RoleID       uuid.UUID `db:"role_id" json:"role_id" validate:"required,uuid"`
	PermissionID uuid.UUID `db:"permission_id" json:"permission_id" validate:"required,uuid"`
}

type UserRole struct {
	UserID uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	RoleID uuid.UUID `db:"role_id" json:"role_id" validate:"required,uuid"`
}
