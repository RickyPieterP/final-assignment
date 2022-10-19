package request

type JWTToken struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
