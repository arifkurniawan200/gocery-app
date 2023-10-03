package model

import "time"

type User struct {
	ID           int64     `json:"id,omitempty" db:"id"`
	Username     string    `json:"username,omitempty" db:"username"`
	PasswordHash string    `json:"password_hash,omitempty" db:"password_hash"`
	IsAdmin      bool      `json:"is_admin,omitempty" db:"is_admin"`
	CreatedAt    time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
