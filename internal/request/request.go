package request

type CreateRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type CreatePreviewRequest struct {
	ID        uint64 `json:"id,omitempty"` // Opcional, depende de cómo estés manejando las IDs
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}
