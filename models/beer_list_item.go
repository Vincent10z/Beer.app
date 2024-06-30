package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (b *BeerListItem) BeforeCreate(db *gorm.DB) error {
	b.ID = utils.GenerateBreweryID()
	return nil
}

type BeerListItem struct {
	ID      string    `gorm:"primaryKey;type:text" json:"id"`
	ListID  string    `gorm:"type:text" json:"list_id"`
	List    BeerList  `gorm:"foreignKey:ListID" json:"list"`
	BeerID  string    `gorm:"type:text" json:"beer_id"`
	Beer    Beer      `gorm:"foreignKey:BeerID" json:"beer"`
	AddedAt time.Time `gorm:"autoCreateTime" json:"added_at"`
}
