package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uuid.UUID `gorm:"primaryKey,type:uuid;default:uuid_generate_v4()"`
	Username     string    `gorm:"uniqueIndex;not null"`
	Mail         string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
}
