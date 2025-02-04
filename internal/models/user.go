package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UpiHandle string    `gorm:"unique;not null;type:varchar(100)"`
	Email     string    `gorm:"unique;not null;type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"not null;type:varchar(100)"`
	Wallets   []Wallet
}
