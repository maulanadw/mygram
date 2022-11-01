package params

import "time"

type CreateCommentRequest struct {
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
}

type CreateCommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetCommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
	Photo     Photo     `json:"photo"`
}

type UpdateCommentRequest struct {
	Message string `json:"message"`
}

type UpdateCommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteCommentResponse struct {
	Message string `json:"message"`
}
