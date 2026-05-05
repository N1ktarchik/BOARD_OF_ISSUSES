package middleware

import (
	"N1ktarchik/Board_of_issues/internal/core/domain"
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"time"
)

type AuthService interface {
	CreateJWT(userID string) (string, error)
	GetUserIdFromJWT(JWT string) (string, error)
	ValidateJWT(JWT string) (*domain.Claims, error)
}

type Cache interface {
	Get(ctx context.Context, key string) (*domain.IdempotencyRecord, error)
	Set(ctx context.Context, key string, data *domain.IdempotencyRecord, TTL time.Duration) error
}

type MiddleWare struct {
	authService AuthService
	cache       Cache
	log         *slog.Logger
}

func NewMiddleWare(authService AuthService, cache Cache, log *slog.Logger) *MiddleWare {
	return &MiddleWare{
		authService: authService,
		cache:       cache,
		log:         log,
	}
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func (rec *responseRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *responseRecorder) Write(b []byte) (int, error) {
	rec.body.Write(b)
	return rec.ResponseWriter.Write(b)
}
