package auth

import (
	"os"
	"time"
)

const (
	secretJWTkey string        = "secretJWT"
	tokenLive    time.Duration = 1 * time.Hour
	autorOfToken string        = "board_of_issuses_app"
)

type ConfigJWT struct {
	SecretKey []byte
	TokenLive time.Duration
	Autor     string
}

func LoadJwtConfig() *ConfigJWT {
	return &ConfigJWT{
		SecretKey: []byte(getEnvData(secretJWTkey)),
		TokenLive: tokenLive,
		Autor:     autorOfToken,
	}
}

func getEnvData(key string) string {
	return os.Getenv(key)
}
