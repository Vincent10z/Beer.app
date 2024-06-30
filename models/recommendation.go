package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (b *Recommendation) BeforeCreate(db *gorm.DB) error {
	b.ID = utils.GenerateBreweryID()
	return nil
}

type Recommendation struct {
	ID            string    `gorm:"primaryKey;type:text" json:"id"`
	UserID        string    `gorm:"type:text" json:"user_id"`
	User          User      `gorm:"foreignKey:UserID" json:"user"`
	BeerID        string    `gorm:"type:text" json:"beer_id"`
	Beer          Beer      `gorm:"foreignKey:BeerID" json:"beer"`
	RecommendedBy string    `gorm:"type:text" json:"recommended_by"`
	Recommender   User      `gorm:"foreignKey:RecommendedBy" json:"recommender"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
}
