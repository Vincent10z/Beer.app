package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (b *Beer) BeforeCreate(db *gorm.DB) error {
	b.ID = utils.GenerateBeerID()
	return nil
}

type Beer struct {
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	ID          string    `gorm:"primaryKey;type:text" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Type        string    `json:"type"`
	StyleID     string    `gorm:"type:text" json:"style_id"`
	Style       BeerStyle `gorm:"foreignKey:StyleID" json:"style"`
	ABV         float32   `json:"abv"`
	IBU         float32   `json:"ibu"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	BreweryID   string    `gorm:"type:text" json:"brewery_id"`
	Brewery     Brewery   `gorm:"foreignKey:BreweryID" json:"brewery"`
}
