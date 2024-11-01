package student

import (
	"github.com/ZYallers/fine/frame/f"
)

type InfoReq struct {
	f.Meta `path:"school/student/info" method:"post"`
	Id     uint `form:"id" validate:"required,min=1" json:"id"` // student id
}

type InfoRes struct {
	Id   uint   `json:"id"`   // student id
	Name string `json:"name"` // student name
}
