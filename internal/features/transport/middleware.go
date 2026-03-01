package tarnsport

import (
	auth "Board_of_issuses/internal/core/auth/jwt"
	"Board_of_issuses/internal/features/transport/http/handlers"
	"context"
	"net/http"
	"strings"

	er "Board_of_issuses/internal/core"
)

type AuthHandler struct {
	JWTManager *auth.JWTService
}

func CreateAuthHandler(JWTManager *auth.JWTService) *AuthHandler {
	return &AuthHandler{
		JWTManager: JWTManager,
	}
}

func (a *AuthHandler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			handlers.RespondWithError(w, http.StatusUnauthorized, "you have been unaothorized. Please login.")
			return
		}

		parts := strings.Split(header, " ")

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			handlers.RespondWithError(w, http.StatusBadRequest, "token is bad data")
			return
		}

		userID, err := a.JWTManager.GetUserIdFromJWT(parts[1])

		if err != nil {
			switch {

			case er.IsError(err, "TOKEN_NOT_VALID"):
				var appErr *er.ErrorApp = err.(*er.ErrorApp)
				handlers.RespondWithError(w, http.StatusUnauthorized, appErr.Message)

			case er.IsError(err, "JWT_METHOD_NOT_VALID"):
				var appErr *er.ErrorApp = err.(*er.ErrorApp)
				handlers.RespondWithError(w, http.StatusBadRequest, appErr.Message)

			default:
				handlers.RespondWithError(w, http.StatusInternalServerError, "error to validate token.Please login.")
			}

			return

		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
