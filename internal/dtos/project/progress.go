package project

type ProgressResponse struct {
	ProjectID      int64   `json:"project_id"`
	TotalTasks     int     `json:"total_tasks"`
	CompletedTasks int     `json:"completed_tasks"`
	Progress       float64 `json:"progress"`
}
