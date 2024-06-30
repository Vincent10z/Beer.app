package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (b *UserActivity) BeforeCreate(db *gorm.DB) error {
	b.ID = utils.GenerateBreweryID()
	return nil
}

type UserActivity struct {
	ID           string    `gorm:"primaryKey;type:text" json:"id"`
	UserID       string    `gorm:"type:text" json:"user_id"`
	User         User      `gorm:"foreignKey:UserID" json:"user"`
	ActivityType string    `gorm:"not null" json:"activity_type"`
	ActivityID   string    `gorm:"type:text" json:"activity_id"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
