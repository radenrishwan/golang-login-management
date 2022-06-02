package domain

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      Role   `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)
