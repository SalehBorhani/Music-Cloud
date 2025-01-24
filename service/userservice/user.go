package userservice

import (
	"github.com/yazdanbhd/Music-Cloud/entity"
	"github.com/yazdanbhd/Music-Cloud/params"
	"github.com/yazdanbhd/Music-Cloud/service/authservice"
	"github.com/yazdanbhd/Music-Cloud/service/totpservice"
)

type Repository interface {
	Register(u entity.User) (entity.User, error)
	IsAuthenticated(userName, password string) (bool, error)
	GetUserID(userName string) (uint, error)
}

type Service struct {
	auth authservice.Service
	totp totpservice.Service
	repo Repository
}

func New(r Repository, auth authservice.Service) Service {
	return Service{repo: r, auth: auth}
}

func (s *Service) UserRegister(req params.RegisterRequest) (params.RegisterResponse, error) {
	// Store the user data to the database
	user := entity.User{
		ID:       0,
		Password: req.Password,
		Email:    req.Email,
		Name:     req.Name,
		UserName: req.UserName,
	}
	u, err := s.repo.Register(user)
	if err != nil {
		return params.RegisterResponse{}, err
	}
	return params.RegisterResponse{UserID: u.ID, UserName: u.UserName, TOTPUri: s.totp.GenerateOTP(user.Email, totpservice.RandomSecret)}, nil
}

func (s *Service) UserLogin(loginReq params.LoginRequest) (params.LoginResponse, error) {
	isAuth, err := s.repo.IsAuthenticated(loginReq.UserName, loginReq.Password)

	if err != nil || isAuth == false {
		return params.LoginResponse{}, err
	}

	userID, err := s.repo.GetUserID(loginReq.UserName)
	if err != nil {
		return params.LoginResponse{}, err
	}

	accessToken, err := s.auth.CreateAccessToken(userID)
	if err != nil {
		return params.LoginResponse{}, err
	}

	refreshToken, err := s.auth.CreateRefreshToken(userID)

	if err != nil {
		return params.LoginResponse{}, err
	}

	return params.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken, UserInfo: struct {
		UserID uint `json:"id"`
	}{UserID: userID}}, nil

}
