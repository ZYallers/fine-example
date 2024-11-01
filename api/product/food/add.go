package food

import (
	"github.com/ZYallers/fine/frame/f"
)

type AddReq struct {
	f.Meta   `path:"product/food/add" method:"post"`
	Name     string `form:"name" validate:"required,max=20" json:"name"`        // food name
	Quantity uint   `form:"quantity" validate:"required,min=1" json:"quantity"` // quantity
}

type AddRes struct {
	Id int `json:"id"` // food id
}
