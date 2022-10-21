package request

type CreateCommentReq struct {
	UserIdComment int    `json:"user_id"`
	PhotoId       int    `json:"photo_id"`
	Tittle        string `json:"tittle"`
	Message       string `json:"message"`
}
