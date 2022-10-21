package response

type UpdateUserRes struct {
	Age       int    `json:"age"`
	Email     string `json:"email"`
	ID        int    `json:"id"`
	Username  string `json:"username"`
	UpdatedAt string `json:"updated_at"`
}
