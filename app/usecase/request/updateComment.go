package request

type UpdateCommentReq struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Message string `json:"message"`
}