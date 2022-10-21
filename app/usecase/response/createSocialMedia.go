package response

type CreateSocialMediaRes struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         int    `json:"user_id"`
	CreatedAt      string `json:"created_at"`
}
