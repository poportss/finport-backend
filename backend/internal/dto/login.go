package dto

type Login struct {
	Provider string `json:"provider"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
