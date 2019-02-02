package users

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User - structure to describe user
type User struct {
	gorm.Model
	Username string    `json:"username,omitempty"`
	Email    string    `json:"email,omitempty"`
	Created  time.Time `json:"created,omitempty"`
	UserID   string    `json:"user_id,omitempty"`
}
