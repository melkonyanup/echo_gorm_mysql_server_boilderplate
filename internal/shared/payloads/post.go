package payloads

// UpdatePost body
type UpdatePostPayload struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}
