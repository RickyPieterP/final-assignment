package request

type DeleteCommentReq struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
}