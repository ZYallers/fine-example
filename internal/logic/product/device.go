package product

import (
	"github.com/ZYallers/fine-example/api/product/device"
	"github.com/ZYallers/fine-example/internal/model"
	"github.com/ZYallers/fine-example/internal/service"
)

type sDevice struct{}

func init() {
	service.RegisterProductDevice(NewDevice())
}

func NewDevice() *sDevice {
	s := &sDevice{}
	return s
}

// Add New device add
func (s *sDevice) Add(req *device.AddReq) (err error) {
	return
}

// Detail Get device detail
func (s *sDevice) Detail(id int) (res model.DeviceDetail) {
	return
}
