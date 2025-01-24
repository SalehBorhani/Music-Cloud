package params

type RegisterRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	UserName string `json:"user_name"`
	UserID   uint   `json:"user_id"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
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
