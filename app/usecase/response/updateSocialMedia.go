package response

import "time"

type UpdateSocialMediaRes struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	UpdateAt       time.Time `json:"updated_at"`
}
