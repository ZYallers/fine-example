package student

import "github.com/ZYallers/fine/frame/f"

type AddReq struct {
	f.Meta `path:"school/student/add" method:"post"`
	Age    uint   `form:"age" validate:"omitempty,min=1,max=100" json:"age"` // age
	Name   string `form:"name" validate:"required" json:"name"`              // name
}

type AddRes struct {
	Id int `json:"id"` // id
}
