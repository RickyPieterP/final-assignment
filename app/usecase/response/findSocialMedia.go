package response

import "time"

type FindSocialMediaRes struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
<<<<<<< HEAD
	SocialMediaURL string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
=======
	SocialMediaURL int       `json:"social_media_url"`
	UserID         string    `json:"user_id"`
>>>>>>> a5220b0de19a6db2e8b8f0d855c81d140b1319dd
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           User      `json:"user"`
}

type User struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	ProfileImageURL int    `json:"profile_image_url"`
}
