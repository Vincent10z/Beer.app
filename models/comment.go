package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (b *Comment) BeforeCreate(db *gorm.DB) error {
	b.ID = utils.GenerateBreweryID()
	return nil
}

type Comment struct {
	ID         string    `gorm:"primaryKey;type:text" json:"id"`
	UserID     string    `gorm:"type:text" json:"user_id"`
	User       User      `gorm:"foreignKey:UserID" json:"user"`
	ReviewID   string    `gorm:"type:text" json:"review_id"`
	ReviewType string    `gorm:"check:review_type IN ('beer', 'brewery')" json:"review_type"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
