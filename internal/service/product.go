// ================================================================================
// Code generated and maintained by Fine CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// Created at 2024/11/01 10:47:12.286.
// ================================================================================

package service

import (
	"github.com/ZYallers/fine-example/api/product/device"
	"github.com/ZYallers/fine-example/api/product/food"
	"github.com/ZYallers/fine-example/internal/model"
)

type (
	IProductFood interface {
		Add(req *food.AddReq) (err error)                                  // Add New food add
		Delete(id int) (err error)                                         // Delete Food delete
		PageList(req *food.PageListReq) (res *food.PageListRes, err error) // PageList Page get food list
	}
	IProductDevice interface {
		Add(req *device.AddReq) (err error)     // Add New device add
		Detail(id int) (res model.DeviceDetail) // Detail Get device detail
	}
)

var (
	localProductFood   IProductFood
	localProductDevice IProductDevice
)

func ProductFood() IProductFood {
	if localProductFood == nil {
		panic("implement not found for interface IProductFood, forgot register?")
	}
	return localProductFood
}

func RegisterProductFood(i IProductFood) {
	localProductFood = i
}

func ProductDevice() IProductDevice {
	if localProductDevice == nil {
		panic("implement not found for interface IProductDevice, forgot register?")
	}
	return localProductDevice
}

func RegisterProductDevice(i IProductDevice) {
	localProductDevice = i
}
