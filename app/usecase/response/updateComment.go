package response

import "time"

type UpdateCommentResp struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	UserId    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
