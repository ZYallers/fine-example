// =================================================================================
// This is auto-generated by Fine CLI tool only once. Fill this file as you wish.
// Created at 2024/10/30 17:58:01.969.
// =================================================================================

package dao

import (
	"github.com/ZYallers/fine-example/internal/dao/internal"
)

// internalDeviceDao is internal type for wrapping internal dao implements.
type internalDeviceDao = *internal.DeviceDao

// deviceDao is the data access object for table device.
// You can define custom methods on it to extend its functionality as you wish.
type deviceDao struct {
	internalDeviceDao
}

var (
	// Device is globally public accessible object for table device operations.
	Device = deviceDao{
		internal.NewDeviceDao(),
	}
)
