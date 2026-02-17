package auth

import (
	"os"
	"time"
)

const (
	secretJWTkey string        = "secretJWT"
	tokenLive    time.Duration = 1 * time.Hour
	autorOfToken string        = "AutorToken"

	defaultAutor string = "board_of_issuses_app"
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
		Autor:     getEnvData(autorOfToken),
	}
}

func getEnvData(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	switch key {
	case autorOfToken:
		return defaultAutor
	default:
		return ""
	}
}
