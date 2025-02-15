package models

import "github.com/google/uuid"

type Wallet struct {
	WalletId   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserUID    uuid.UUID `gorm:"type:uuid;not null"`
	User       User      `gorm:"foreignKey:UserUID;constraint:OnDelete:CASCADE;"`
	Address    string    `gorm:"not null;"`
	PrivateKey string    `gorm:"not null;"`
	IsPrimary  bool      `gorm:"default:false"`
	Chain      string    `gorm:"not null;type:varchar(10)"`
}
