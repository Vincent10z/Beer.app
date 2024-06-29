package models

type Brewery struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	City     string    `json:"city"`
	State    string    `json:"state"`
	ZipCode  string    `json:"zip_code"`
	Products []Product `json:"products,omitempty"`
}
