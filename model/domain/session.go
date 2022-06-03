package domain

type Session struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	ExpiredAt string `json:"expired_at"`
}
