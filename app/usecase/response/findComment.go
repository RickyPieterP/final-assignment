package response

import "time"

type FindCommentResp struct {
	Id        int         `json:"id"`
	Message   string      `json:"message"`
	PhotoId   int         `json:"photo_id"`
	UserId    int         `json:"user_id"`
	CreatedAt time.Time   `json:"created_date"`
	UpdatedAt time.Time   `json:"updated_at"`
	User      UserComment `json:"User"`
}

type UserComment struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
