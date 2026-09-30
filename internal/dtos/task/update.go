package task

type UpdateRequest struct {
	Title    string `json:"title"`
	MemberID int64  `json:"member_id"`
	DueDate  string `json:"due_date"`
}
