package uuid

import (
	"github.com/satori/go.uuid"
)

// GetUUID gives the uuid
func GetUUID() string {
	return uuid.NewV4().String()
}

