package domain

type User struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	IdentityNumber string `json:"identity_number"`
}
