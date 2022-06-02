package web

type DeleteUserRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}
