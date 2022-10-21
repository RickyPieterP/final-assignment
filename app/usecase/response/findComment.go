package response

import "time"

type FindCommentResp struct {
	Id        int          `json:"id"`
	Message   string       `json:"message"`
	PhotoId   int          `json:"photo_id"`
	UserId    int          `json:"user_id"`
	UpdatedAt time.Time    `json:"updated_at"`
	CreatedAt time.Time    `json:"created_at"`
	User      CommentUser  `json:"User"`
	Photo     CommentPhoto `json:"Photo"`
}

type CommentUser struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CommentPhoto struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}
