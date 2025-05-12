package util

import (
	"github.com/google/uuid"
)

func GenerateRandomAPIKey() string {
	return uuid.NewString()
}
