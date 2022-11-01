package params

import "time"

type User struct {
	ID       *int   `json:"id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegisterUserRequest struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type RegisterUserResponse struct {
	ID       int    `json:"id"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UpdateUserResponse struct {
	ID        int       `json:"id"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
