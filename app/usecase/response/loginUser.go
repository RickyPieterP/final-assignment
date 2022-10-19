package response

type LoginUserResp struct {
	Token   string `json:"token,omitempty"`
	Message string `json:"message,omitempty"`
}
