package models

import (
	"time"
)

//type Beer struct {
//	CreatedAt   string `json:"createdAt"`
//	UpdatedAt   string `json:"updatedAt"`
//	Id          string `json:"id"`
//	Name        string `json:"name"`
//	Type        string `json:"type"`
//	Grains      string `json:"grains"`
//	ABV         int    `json:"abv"`
//	Price       int    `json:"price"`
//	Rating      int    `json:"rating"`
//	Quantity    int    `json:"quantity"`
//	IsAvailable bool   `json:"isAvailable"`
//	IsDeleted   bool   `json:"isDeleted"`
//}

type Beer struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	BreweryID   uint      `json:"breweryId"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	ABV         int       `json:"abv"`
}
