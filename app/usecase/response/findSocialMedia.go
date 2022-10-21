package response

import "time"

type FindSocialMediaRes struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           User      `json:"user"`
}

type User struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	ProfileImageURL int    `json:"profile_image_url"`
}
