package dto

type Register_RequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
