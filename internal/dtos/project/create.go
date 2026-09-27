package project

type CreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     int64  `json:"owner_id"`
}
