package product

import (
	"github.com/ZYallers/fine-example/api/product/food"
	"github.com/ZYallers/fine-example/internal/service"
)

type sFood struct{}

func init() {
	service.RegisterProductFood(NewFood())
}

func NewFood() *sFood {
	s := &sFood{}
	return s
}

// Add New food add
func (s *sFood) Add(req *food.AddReq) (err error) {
	return
}

// Delete Food delete
func (s *sFood) Delete(id int) (err error) {
	return
}

// PageList Page get food list
func (s *sFood) PageList(req *food.PageListReq) (res *food.PageListRes, err error) {
	return
}
