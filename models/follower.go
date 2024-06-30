package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (b *Follower) BeforeCreate(db *gorm.DB) error {
	b.ID = utils.GenerateBreweryID()
	return nil
}

type Follower struct {
	ID         string    `gorm:"primaryKey;type:text" json:"id"`
	FollowerID string    `gorm:"type:text" json:"follower_id"`
	Follower   User      `gorm:"foreignKey:FollowerID" json:"follower"`
	FollowedID string    `gorm:"type:text" json:"followed_id"`
	Followed   User      `gorm:"foreignKey:FollowedID" json:"followed"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}
