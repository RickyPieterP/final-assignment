package request

type CreatePhotoReq struct {
	Title string `json:"title" validate:"required"`
	Caption string `json:"caption" validate:"required"`
	PhotoUrl string `json:"photo_url" validate:"required"`
	UserId int `json:"user_id" validate:"required"`
}