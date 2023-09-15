package models

import (
	"time"

	"github.com/google/uuid"
)

// Users struct to describe users object.
type Users struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Username     string    `db:"username" json:"username" validate:"required,lte=50"`
	Password     string    `db:"password" json:"password" validate:"required,lte=1000"`
	ActiveStatus int       `db:"active_status" json:"active_status" validate:"required,len=1"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

// Users struct to describe users object.
type Profile struct {
	ID             uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	UserId         uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Email          string    `db:"email" json:"email" validate:"lte=50"`
	Role           string    `db:"role" json:"role" validate:"required,lte=50"`
	FullName       string    `db:"full_name" json:"full_name" validate:"lte=100"`
	VerifiedStatus int       `db:"verified_status" json:"verified_status" validate:"required,len=1"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

type SignUpInput struct {
	FullName        string `json:"full_name" validate:"required,lte=100"`
	Email           string `json:"email" validate:"required,min=8"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}
