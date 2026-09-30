package domain

import "time"

type User struct {
	ID        int64     `json:"id"`
	CivilID   string    `json:"civil_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
