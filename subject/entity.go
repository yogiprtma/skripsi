package subject

import "time"

type Subject struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
