package request

type CreateCommentReq struct {
	UserId  int    `json:"user"`
	PhotoId int    `json:"photo_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}
