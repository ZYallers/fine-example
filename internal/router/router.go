package router

import (
	"github.com/ZYallers/fine-example/internal/controller/product/device"
	"github.com/ZYallers/fine-example/internal/controller/product/food"
	"github.com/ZYallers/fine-example/internal/controller/school/student"
	"github.com/ZYallers/fine/frame/frouter"
)

var Router = frouter.Router{}

func init() {
	Router.Register(
		device.New(),
		food.New(),
		student.New(),
	)
}
