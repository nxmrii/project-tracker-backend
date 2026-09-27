package member

type ContributionResponse struct {
	MemberID       int64   `db:"member_id" json:"member_id"`
	MemberName     string  `db:"member_name" json:"member_name"`
	TotalTasks     int     `db:"total_tasks" json:"total_tasks"`
	CompletedTasks int     `db:"completed_tasks" json:"completed_tasks"`
	Contribution   float64 `db:"contribution" json:"contribution"`
}
