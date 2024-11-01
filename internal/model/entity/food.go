// =================================================================================
// Code generated and maintained by Fine CLI tool. DO NOT EDIT.
// Created at 2024/10/30 17:57:57.902.
// =================================================================================

package entity

// Food is the golang structure for table food.
type Food struct {
	Id       uint64  `json:"id"       gorm:"column:id;type:bigint(20) unsigned;not null;primaryKey;autoIncrement"  ` // Primary ID
	Name     string  `json:"name"     gorm:"column:name;type:varchar(30);not null;default:''"                      ` // Food Name
	Price    float64 `json:"price"    gorm:"column:price;type:double(11,2);not null;default:0.00"                  ` // Price
	Quantity uint    `json:"quantity" gorm:"column:quantity;type:int(11) unsigned;not null;default:0"              ` // Quantity
}
