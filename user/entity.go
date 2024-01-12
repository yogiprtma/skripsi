package user

import "time"

type User struct {
	ID           int
	Name         string
	NPM          string
	Department   string
	PasswordHash string
	PublicKey    string
	PrivateKey   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
