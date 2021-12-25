// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewUser struct {
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type RefreshTokenData struct {
	Token string `json:"token"`
}

type TokenData struct {
	Role     string `json:"role"`
	Username string `json:"username"`
}

type User struct {
	ID          string `json:"id"`
	Role        string `json:"role"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	LastLoginAt string `json:"lastLoginAt"`
}
