package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
)

func (b *BeerStyle) BeforeCreate(tx *gorm.DB) error {
	b.ID = utils.GenerateBreweryID()
	return nil
}

type BeerStyle struct {
	ID          string `gorm:"primaryKey;type:text" json:"id"`
	Name        string `gorm:"uniqueIndex;not null" json:"name"`
	Description string `json:"description"`
}
