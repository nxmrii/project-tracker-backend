package domain

import "time"

type Project struct { //sqlx =>  go : postgresSQL
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Deadline  time.Time `db:"deadline" json:"deadline"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
