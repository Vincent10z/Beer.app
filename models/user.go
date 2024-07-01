package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.ID = utils.GenerateUserID()
	return nil
}

type User struct {
	ID             string    `gorm:"primaryKey;type:text" json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Username       string    `gorm:"uniqueIndex;not null" json:"username"`
	Email          string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash   string    `gorm:"not null" json:"-"`
	ProfilePicture string    `json:"profile_picture"`
	Bio            string    `json:"bio"`
	Address        string    `json:"address"`
	City           string    `json:"city"`
	ZipCode        string    `json:"zip_code"`
	State          string    `json:"state"`
	Country        string    `json:"country"`
	Phone          string    `json:"phone"`
	Role           string    `gorm:"default:user" json:"role"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Account        *Account  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"account"`
}
