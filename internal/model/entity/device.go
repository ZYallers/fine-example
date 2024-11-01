// =================================================================================
// Code generated and maintained by Fine CLI tool. DO NOT EDIT.
// Created at 2024/10/30 17:58:02.004.
// =================================================================================

package entity

// Device is the golang structure for table device.
type Device struct {
	Id       uint64  `json:"id"        gorm:"column:id;type:bigint(20) unsigned;not null;primaryKey;autoIncrement"  ` // Primary ID
	DeviceId string  `json:"device_id" gorm:"column:device_id;type:varchar(32);not null;default:''"                 ` // Device ID
	Name     string  `json:"name"      gorm:"column:name;type:varchar(30);not null;default:''"                      ` // Device Name
	Price    float64 `json:"price"     gorm:"column:price;type:double(11,2);not null;default:0.00"                  ` // Price
	Quantity uint    `json:"quantity"  gorm:"column:quantity;type:int(11) unsigned;not null;default:0"              ` // Quantity
}
