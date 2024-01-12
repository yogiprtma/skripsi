package employee

import (
	"time"
)

type Employee struct {
	ID           int
	Nip          string
	Name         string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
