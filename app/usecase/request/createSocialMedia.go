package request

type CreateSocialMediaReq struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}
