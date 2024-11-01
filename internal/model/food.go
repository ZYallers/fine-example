package model

type FoodPageList struct {
	Id       int    `json:"id"`       // Food ID
	Name     string `json:"name"`     // Food Name
	Quantity int    `json:"quantity"` // Food Quantity
}
