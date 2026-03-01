package authjwt

import auth "Board_of_issuses/internal/core/auth/jwt"

type AuthManager struct {
	JWTManager *auth.JWTService
}

func CreateAuthManager(JWTManager *auth.JWTService) *AuthManager {
	return &AuthManager{
		JWTManager: JWTManager,
	}
}
