// ==========================================================================
// Code generated and maintained by Fine CLI tool. DO NOT EDIT.
// Created at 2024/10/30 17:58:01.972.
// ==========================================================================

package internal

import (
	"github.com/ZYallers/fine/database/fmysql"
	"gorm.io/gorm"
)

// DeviceDao is the data access object for table device.
type DeviceDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns DeviceColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceColumns defines and stores column names for table device.
type DeviceColumns struct {
	Id       string // primary ID
	DeviceId string // device ID
	Name     string // device name
	Price    string // price
	Quantity string // quantity
}

// deviceColumns holds the columns for table device.
var deviceColumns = DeviceColumns{
	Id:       "id",
	DeviceId: "device_id",
	Name:     "name",
	Price:    "price",
	Quantity: "quantity",
}

// NewDeviceDao creates and returns a new dao object for table data access.
func NewDeviceDao() *DeviceDao {
	return &DeviceDao{
		group:   "test2",
		table:   "device",
		columns: deviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current dao.
func (dao *DeviceDao) DB() *gorm.DB { return fmysql.DB(dao.group) }

// Table returns the table name of current dao.
func (dao *DeviceDao) Table() string { return dao.table }

// Columns returns all column names of current dao.
func (dao *DeviceDao) Columns() DeviceColumns { return dao.columns }

// Group returns the configuration group name of database of current dao.
func (dao *DeviceDao) Group() string { return dao.group }

// Session creates and returns the session for current dao.
func (dao *DeviceDao) Session() *gorm.DB { return dao.DB().Table(dao.table) }

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not commit or rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceDao) Transaction(f func(tx *gorm.DB) error) error { return dao.DB().Transaction(f) }
