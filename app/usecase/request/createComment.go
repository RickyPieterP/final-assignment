package request

type CreateCommentReq struct {
	UserIdComment int    `json:"user_id" validate:"required"`
	PhotoId       int    `json:"photo_id" validate:"required"`
	Message       string `json:"message" validate:"required"`
}
