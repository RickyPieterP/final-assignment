package response

<<<<<<< HEAD
import "time"

type CreateSocialMediaRes struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
=======
type CreateSocialMediaRes struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         int    `json:"user_id"`
	CreatedAt      string `json:"created_at"`
>>>>>>> a5220b0de19a6db2e8b8f0d855c81d140b1319dd
}
