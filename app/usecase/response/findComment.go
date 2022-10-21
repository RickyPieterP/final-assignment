package response

import "time"

type FindCommentResp struct {
	Id        int          `json:"id"`
	Message   string       `json:"message"`
	PhotoId   int          `json:"photo_id"`
	UserId    int          `json:"user_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      UserComment  `json:"User"`
	Photo     PhotoComment `json:"Photo"`
}

type UserComment struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoComment struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}
