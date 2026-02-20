package auth

import (
	dn "Board_of_issuses/internal/core/domains"
)

type Auth interface {
	CreateJwt(userID int, email string) (string, error)
	ValidateJWT(jwtToken string) (*dn.Claims, error)
}
