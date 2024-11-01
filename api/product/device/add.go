package device

import (
	"github.com/ZYallers/fine/frame/f"
)

type AddReq struct {
	f.Meta   `path:"product/device/add" method:"post"`
	Name     string  `form:"name" validate:"required" json:"name"`               // device name
	Quantity uint    `form:"quantity" validate:"required,min=1" json:"quantity"` // device quantity
	Price    float64 `form:"price" validate:"required,min=0.01" json:"price"`    // price
}

type AddRes struct {
	Id   int         `json:"id"`   // device ID
	Name string      `json:"name"` // device name
	Now  string      `json:"now"`  // now time
	Req  interface{} `json:"req"`  // request params
}
