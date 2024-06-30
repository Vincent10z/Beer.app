package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (b *Checkin) BeforeCreate(db *gorm.DB) error {
	b.ID = utils.GenerateBreweryID()
	return nil
}

type Checkin struct {
	ID        string    `gorm:"primaryKey;type:text" json:"id"`
	UserID    string    `gorm:"type:text" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	BreweryID string    `gorm:"type:text" json:"brewery_id"`
	Brewery   Brewery   `gorm:"foreignKey:BreweryID" json:"brewery"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
