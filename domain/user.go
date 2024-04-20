package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `json:"id" gorm:"type:BINARY;default:(UUID_TO_BIN(UUID()));primaryKey;not null"`
	Username  string    `json:"username" gorm:"not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime; not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime; not null"`
}
