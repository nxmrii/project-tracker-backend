package member

// this is the box that response frontend data
type CreateRequest struct {
	Name string `json:"name"`
	Role string `json:"role"`
}
