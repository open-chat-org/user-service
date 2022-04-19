package model

type UserModel struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Profile  string `json:"profile"`
	ID       int    `json:"id"`
}
