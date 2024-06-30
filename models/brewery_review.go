package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (b *BreweryReview) BeforeCreate(db *gorm.DB) error {
	b.ID = utils.GenerateBreweryID()
	return nil
}

type BreweryReview struct {
	ID        string    `gorm:"primaryKey;type:text" json:"id"`
	BreweryID string    `gorm:"type:text" json:"brewery_id"`
	Brewery   Brewery   `gorm:"foreignKey:BreweryID" json:"brewery"`
	UserID    string    `gorm:"type:text" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Rating    int       `gorm:"check:rating >= 0 AND rating <= 5" json:"rating"`
	Comment   string    `json:"comment"`
	Likes     int       `gorm:"default:0" json:"likes"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
