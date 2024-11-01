// =================================================================================
// Code generated and maintained by Fine CLI tool. DO NOT EDIT.
// Created at 2024/10/30 18:36:15.428.
// =================================================================================

package product

import (
	"github.com/gin-gonic/gin"

	"github.com/ZYallers/fine-example/api/product/device"
	"github.com/ZYallers/fine-example/api/product/food"
)

type IProductDevice interface {
	Add(ctx *gin.Context, req *device.AddReq) (res *device.AddRes, err error)
	Detail(ctx *gin.Context, req *device.DetailReq) (res *device.DetailRes, err error)
}

type IProductFood interface {
	Add(ctx *gin.Context, req *food.AddReq) (res *food.AddRes, err error)
	PageList(ctx *gin.Context, req *food.PageListReq) (res *food.PageListRes, err error)
}
