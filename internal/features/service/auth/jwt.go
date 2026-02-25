package auth

import (
	"log"
	"time"

	er "Board_of_issuses/internal/core"

	dn "Board_of_issuses/internal/core/domains"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	config *ConfigJWT
}

func CreateJWTService(cfg *ConfigJWT) *JWTService {
	return &JWTService{
		config: cfg,
	}
}

func (s *JWTService) CreateJwt(userID int, email string) (string, error) {
	secret := s.config.SecretKey

	if len(secret) == 0 {
		log.Panicf("env key not faund")
	}

	claims := dn.Claims{
		UserId: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.config.TokenLive)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    s.config.Autor,
		},
	}

	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (s *JWTService) ValidateJWT(jwtToken string) (*dn.Claims, error) {

	claims := &dn.Claims{}

	token, err := jwt.ParseWithClaims(jwtToken, claims, func(t *jwt.Token) (interface{}, error) {

		if t.Method != jwt.SigningMethodHS256 {
			return nil, er.JWTMethodError()
		}

		return s.config.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return claims, nil
	}

	return nil, er.JWTTokenNotValid()
}
