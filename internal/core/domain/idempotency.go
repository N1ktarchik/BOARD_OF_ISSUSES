package domain

type IdempotencyRecord struct {
	StatusCode int    `json:"status_code"`
	Body       []byte `json:"body"`
}
