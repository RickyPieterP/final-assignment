package response

import "time"

type UpdateCommentResp struct {
	Id        int       `json:"id"`
	Message   string    `json:"message"`
	UserId    int       `json:"user_id"`
	PhotoId   int       `json:"photo_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
