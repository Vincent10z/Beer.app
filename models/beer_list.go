package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (b *BeerList) BeforeCreate(db *gorm.DB) error {
	b.ID = utils.GenerateBreweryID()
	return nil
}

type BeerList struct {
	ID          string    `gorm:"primaryKey;type:text" json:"id"`
	UserID      string    `gorm:"type:text" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	IsPublic    bool      `gorm:"default:false" json:"is_public"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
