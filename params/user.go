package params

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	UserName string `json:"user_name"`
	TOTPUri  string `json:"totp_uri"`
	UserID   uint   `json:"user_id"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	TOTPCode string `json:"totp_code"`
	Password string `json:"password"`
}

type UserInfo struct {
	UserID uint `json:"id"`
}

type LoginResponse struct {
	UserInfo     UserInfo `json:"user_info"`
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
}
