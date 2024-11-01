package food

import (
	"github.com/ZYallers/fine-example/internal/model"
	"github.com/ZYallers/fine-example/util/api"
	"github.com/ZYallers/fine/frame/f"
)

type PageListReq struct {
	f.Meta `path:"product/food/page/list" method:"post"`
	api.Page
	Name string `form:"name" validate:"omitempty,max=20" json:"name"` // food name
}

type PageListRes struct {
	Total uint64               `json:"total"` // total
	List  []model.FoodPageList // list data
}
