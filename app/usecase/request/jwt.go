package request

type JWTToken struct {
	Id int `json:"id"`
	Username string `json:"username"`
}