package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (b *Event) BeforeCreate(db *gorm.DB) error {
	b.ID = utils.GenerateBreweryID()
	return nil
}

type Event struct {
	ID          string    `gorm:"primaryKey;type:text" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	BreweryID   string    `gorm:"type:text" json:"brewery_id"`
	Brewery     Brewery   `gorm:"foreignKey:BreweryID" json:"brewery"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
