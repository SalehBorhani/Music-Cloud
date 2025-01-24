package totpservice

import (
	"github.com/xlzd/gotp"
)

const RandomSecret = "BLISDSSRG4UMCWWQ445CQT46YQ"

type Config struct {
	AppName string
}

type Service struct {
	config Config
}

func New(config Config) Service {
	return Service{config: config}
}

func (s Service) GenerateOTP(email, randomSecret string) string {
	totp := gotp.NewDefaultTOTP(randomSecret)

	uri := totp.ProvisioningUri(email, s.config.AppName)
	return uri
}

func (s Service) ValidateOTP(randomSecret string, otp string) (bool, error) {
	totp := gotp.NewDefaultTOTP(randomSecret)

	if totp.Now() != otp {
		return false, nil
	}
	return true, nil
}
