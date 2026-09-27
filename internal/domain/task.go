package domain

import "time"

type Task struct {
	ID          int64     `db:"id" json:"id"`
	ProjectID   int64     `db:"project_id" json:"project_id"`
	MemberID    int64     `db:"member_id" json:"member_id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Status      string    `db:"status" json:"status"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
