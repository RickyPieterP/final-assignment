package request

type UpdateComment struct {
	Id           int    `json:"id"`
	TitleValue   string `json:"title"`
	MessageValue string `json:"message"`
	UserId       int    `json:"user_id"`
}
