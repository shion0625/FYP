package domain

import (
	"time"
)

type RefreshSession struct {
	TokenID      string    `gorm:"primaryKey;not null"    json:"tokenId"`
	UserID       string    `gorm:"not null"               json:"userId"`
	RefreshToken string    `gorm:"not null"               json:"refreshToken"`
	ExpireAt     time.Time `gorm:"not null"               json:"expireAt"`
	IsBlocked    bool      `gorm:"not null;default:false" json:"isBlocked"`
}
