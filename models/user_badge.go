package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (b *UserBadge) BeforeCreate(db *gorm.DB) error {
	b.ID = utils.GenerateBreweryID()
	return nil
}

type UserBadge struct {
	ID       string    `gorm:"primaryKey;type:text" json:"id"`
	UserID   string    `gorm:"type:text" json:"user_id"`
	User     User      `gorm:"foreignKey:UserID" json:"user"`
	BadgeID  string    `gorm:"type:text" json:"badge_id"`
	Badge    Badge     `gorm:"foreignKey:BadgeID" json:"badge"`
	EarnedAt time.Time `gorm:"autoCreateTime" json:"earned_at"`
}
