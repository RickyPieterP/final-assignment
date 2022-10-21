package request

type UpdateComment struct {
	Id           int    `json:"id"`
	MessageValue string `json:"message" validate:"required"`
	UserId       int    `json:"user_id"`
}
