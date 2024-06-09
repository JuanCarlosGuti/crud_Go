package domain

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}
