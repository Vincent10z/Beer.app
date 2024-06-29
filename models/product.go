package models

type Product struct {
	autodatedcolumns
	CreatedAt   string    `json:"createdAt"`
	UpdatedAt   string    `json:"updatedAt"`
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Grains      string    `json:"grains"`
	Price       int       `json:"price"`
	Rating      int       `json:"rating"`
	Reviews     []*Review `json:"reviews,omitempty"`
	Quantity    int       `json:"quantity"`
	IsAvailable bool      `json:"isAvailable"`
	IsDeleted   bool      `json:"isDeleted"`
}
