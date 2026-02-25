package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	er "Board_of_issuses/internal/core"
)

func (h *UserHandler) AuthMiddleware(nextHandle func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		parts := strings.Split(header, " ")

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		userID, err := h.serv.CreateJWT(ctx, parts[1])

		if err != nil {
			w.Header().Set("Content-Type", "application/json")

			var responseErr []byte
			switch {

			case er.IsError(err, "TOKEN_NOT_VALID"):
				var appErr *er.ErrorApp = err.(*er.ErrorApp)
				responseErr, _ = json.Marshal(appErr.Message)
				w.WriteHeader(http.StatusUnauthorized)

			case er.IsError(err, "JWT_METHOD_NOT_VALID"):
				var appErr *er.ErrorApp = err.(*er.ErrorApp)
				responseErr, _ = json.Marshal(appErr.Message)
				w.WriteHeader(http.StatusBadRequest)

			default:
				w.WriteHeader(http.StatusInternalServerError)
			}

			if len(responseErr) != 0 {
				w.Write(responseErr)
			}

			return

		}

		ctx = context.WithValue(r.Context(), "userID", userID)
		nextHandle(w, r.WithContext(ctx))
	}
}
