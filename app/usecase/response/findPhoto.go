package response

import "time"

type FindPhotoResp struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User UserPhoto `json:"User"`
}

type UserPhoto struct {
	Email string `json:"email"`
	Username string `json:"username"`
}