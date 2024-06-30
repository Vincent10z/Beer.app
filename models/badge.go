package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
)

func (b *Badge) BeforeCreate(db *gorm.DB) error {
	b.ID = utils.GenerateBreweryID()
	return nil
}

type Badge struct {
	ID          string `gorm:"primaryKey;type:text" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
