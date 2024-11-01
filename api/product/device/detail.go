package device

import (
	"github.com/ZYallers/fine/frame/f"
)

type DetailReq struct {
	f.Meta   `path:"product/device/detail" method:"post"`
	Name     string  `form:"name" validate:"required,max=20" json:"name"`        // device name
	Quantity uint    `form:"quantity" validate:"required,min=1" json:"quantity"` // quantity
	Price    float64 `form:"price" validate:"required,min=0.01" json:"price"`    // price
}

type DetailRes struct {
}
