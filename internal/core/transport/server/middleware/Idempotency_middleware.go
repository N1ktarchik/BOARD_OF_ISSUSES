package middleware

import (
	"N1ktarchik/Board_of_issues/internal/core/domain"
	"N1ktarchik/Board_of_issues/internal/core/errors"
	resp "N1ktarchik/Board_of_issues/internal/core/transport/response"
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"time"
)

func (m *MiddleWare) IdempotencyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			m.log.Debug("idempotency middleware: method is not POST", slog.Any("method", r.Method))
			next.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()

		requestKey := r.Header.Get("X-Req-Key")
		if requestKey == "" {
			m.log.Debug("idempotency middleware: idempotency key not found")
			next.ServeHTTP(w, r)
			return
		}

		userID, ok := domain.GetUserID(ctx)
		if !ok {
			m.log.Warn("idempotency middleware: id not set")
			resp.RespondWithError(w, errors.BadRequest())
			return
		}

		if userID == "" {
			m.log.Warn("idempotency middleware: id not set")
			resp.RespondWithError(w, errors.BadRequest())
			return
		}

		redisKey := "idemp:" + userID + ":" + requestKey

		data, err := m.cache.Get(ctx, redisKey)
		if err == nil && data != nil {
			m.log.Debug("idempotency hit", "key", redisKey)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(data.StatusCode)
			_, _ = w.Write(data.Body)
			return
		}

		rec := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           &bytes.Buffer{},
		}

		next.ServeHTTP(rec, r)

		if rec.statusCode >= 200 && rec.statusCode < 300 {
			data := &domain.IdempotencyRecord{
				StatusCode: rec.statusCode,
				Body:       rec.body.Bytes(),
			}

			redisCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			_ = m.cache.Set(redisCtx, redisKey, data, 24*time.Hour)
		}

	})
}
