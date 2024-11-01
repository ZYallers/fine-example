package model

type DeviceDetail struct {
	Id       int    `json:"id"`       // Device ID
	Name     string `json:"name"`     // Device Name
	Quantity int    `json:"quantity"` // Device Quantity
}
