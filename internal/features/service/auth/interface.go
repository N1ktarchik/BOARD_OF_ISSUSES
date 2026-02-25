package auth

import (
	dn "Board_of_issuses/internal/core/domains"
)

type Auth interface {
	Create(userID int, email string) (string, error)
	Validate(jwtToken string) (*dn.Claims, error)
}
