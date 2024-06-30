package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (b *EventAttendee) BeforeCreate(db *gorm.DB) error {
	b.ID = utils.GenerateBreweryID()
	return nil
}

type EventAttendee struct {
	ID        string    `gorm:"primaryKey;type:text" json:"id"`
	EventID   string    `gorm:"type:text" json:"event_id"`
	Event     Event     `gorm:"foreignKey:EventID" json:"event"`
	UserID    string    `gorm:"type:text" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
