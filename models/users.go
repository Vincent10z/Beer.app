package models

type User struct {
	AccountId int    `json:"AccountId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
