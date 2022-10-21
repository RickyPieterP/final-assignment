package response

import "time"

type CreateCommentResp struct {
	Id        int    `json:"id"`
	Mesage    string `json:"message"`
	PhotoId   int    `json:"photo_id"`
	UserId    int    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}