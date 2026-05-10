package middleware

import (
	"N1ktarchik/Board_of_issues/internal/core/domain"
	"N1ktarchik/Board_of_issues/internal/core/errors"
	resp "N1ktarchik/Board_of_issues/internal/core/transport/response"
	"log/slog"
	"net/http"
	"time"
)

const (
	limit  int           = 15
	window time.Duration = time.Minute
)

func (m *MiddleWare) RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID, ok := domain.GetUserID(ctx)
		if !ok {
			m.log.Warn("rate limit middleware: id not set")
			resp.RespondWithError(w, errors.BadRequest())
			return
		}

		allowed, err := m.cache.RateLimit(ctx, userID, limit, window)
		if err != nil {
			m.log.Error("rate limit middleware: error checking rate limit", "err", err)
			next.ServeHTTP(w, r)
			return
		}

		if !allowed {
			m.log.Debug("rate limit middleware: rate limit exceeded", slog.String("user_id", userID))
			resp.RespondWithError(w, errors.TooManyRequests(userID))
			return
		}

		next.ServeHTTP(w, r)

	})
}
