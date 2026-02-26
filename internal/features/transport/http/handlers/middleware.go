package handlers

import (
	"context"
	"net/http"
	"strings"
	"time"

	er "Board_of_issuses/internal/core"
	respond "Board_of_issuses/internal/features/transport/http"
)

func (h *UserHandler) AuthMiddleware(nextHandle func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			respond.RespondWithError(w, http.StatusUnauthorized, "you have been unaothorized. Please login.")
			return
		}

		parts := strings.Split(header, " ")

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			respond.RespondWithError(w, http.StatusBadRequest, "token is bad data")
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		userID, err := h.serv.CreateJWT(ctx, parts[1])

		if err != nil {
			switch {

			case er.IsError(err, "TOKEN_NOT_VALID"):
				var appErr *er.ErrorApp = err.(*er.ErrorApp)
				respond.RespondWithError(w, http.StatusUnauthorized, appErr.Message)

			case er.IsError(err, "JWT_METHOD_NOT_VALID"):
				var appErr *er.ErrorApp = err.(*er.ErrorApp)
				respond.RespondWithError(w, http.StatusBadRequest, appErr.Message)

			default:
				respond.RespondWithError(w, http.StatusInternalServerError, "error to validate token.Please login.")
			}

			return

		}

		ctx = context.WithValue(r.Context(), "userID", userID)
		nextHandle(w, r.WithContext(ctx))
	}
}
