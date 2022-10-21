package request

type CreateCommentReq struct {
	UserIdComment int    `json:"user_id"`
	PhotoId       int    `json:"photo_id"`
	Message       string `json:"message"`
}
