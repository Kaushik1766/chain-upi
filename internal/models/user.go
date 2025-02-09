package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UpiHandle string    `gorm:"unique;not null;type:varchar(100)"`
	Email     string    `gorm:"unique;not null;type:varchar(100)"`
	Password  string    `gorm:"not null;type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Name      string    `gorm:"not null;type:varchar(100)"`
	Wallets   []Wallet
}
