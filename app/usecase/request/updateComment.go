package request

type UpdateComment struct {
	Id           int    `json:"id"`
	MessageValue string `json:"message"`
	UserId       int    `json:"user_id"`
}
