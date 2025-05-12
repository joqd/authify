package response

import "time"

// ================================== Create Api Key
type CreatedApiKey struct {
	ID        uint      `json:"id"`
	Key       string    `json:"key"`
	Owner     string    `json:"owner"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateApiKeyResponseWrapper struct {
	BaseResponse
	Result CreatedApiKey `json:"result"`
}
