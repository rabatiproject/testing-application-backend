package jwt

type UserCredential struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
