package request

import "time"

// ================================== Create Api Key
type CreateApiKeyRequest struct {
	Owner     string    `json:"owner" validate:"required,min=2,max=30" example:"ackerman.ir"`
	Name      string    `json:"name" validate:"required,min=2,max=30" example:"ackerman1"`
	IsActive  bool      `json:"is_active" validate:"omitempty"`
	ExpiresAt time.Time `json:"expires_at" validate:"required" example:"2025-05-11 14:31:50.523382"`
}
