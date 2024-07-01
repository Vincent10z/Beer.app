package models

import (
	"Beer.app/utils"
	"gorm.io/gorm"
	"time"
)

func (a *Account) BeforeCreate(db *gorm.DB) error {
	a.ID = utils.GenerateAccountID()
	return nil
}

type Account struct {
	ID            string    `gorm:"primaryKey"`
	UserID        string    `gorm:"not null"`
	PlanID        string    `gorm:"size:255"`
	StartDate     time.Time `gorm:"type:date"`
	EndDate       time.Time `gorm:"type:date"`
	Status        string    `gorm:"size:50"`
	PaymentMethod string    `gorm:"size:100"`
}
