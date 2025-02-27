package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (b *Brewery) BeforeCreate(db *gorm.DB) error {
	b.ID = utils.GenerateBreweryID()
	return nil
}

type Brewery struct {
	ID          string    `gorm:"primaryKey;type:text" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	ZipCode     string    `json:"zip_code"`
	Products    []*Beer   `json:"products,omitempty"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	Logo        string    `json:"logo"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
