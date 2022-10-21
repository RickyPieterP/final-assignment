package request

type CreateCommentReq struct {
	UserId  int `json:"user"`
	PhotoId int`json:"photo_id"`
	Message string `json:"messsage"`
}
