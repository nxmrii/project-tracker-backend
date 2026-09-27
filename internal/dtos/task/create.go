package task

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	MemberID    int64  `json:"member_id"`
	Status      string `json:"status"`
}
