package dto

type CreateUser_RequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
