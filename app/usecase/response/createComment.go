package response

import "time"

type CreateCommentResp struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	PhotoId   int       `json:"photo_id"`
	Tittle    string    `json:"tittle"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_date"`
}
