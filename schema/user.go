package schema

type User struct {
	BaseSchema
	Username string `json:"username"`
	Email    string `json:"email"`
}

type TokenBack struct {
	Token string `json:"token"`
}
