// ==========================================================================
// Code generated and maintained by Fine CLI tool. DO NOT EDIT.
// Created at 2024/10/30 17:57:57.869.
// ==========================================================================

package internal

import (
	"github.com/ZYallers/fine/database/fmysql"
	"gorm.io/gorm"
)

// FoodDao is the data access object for table food.
type FoodDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns FoodColumns // columns contains all the column names of Table for convenient usage.
}

// FoodColumns defines and stores column names for table food.
type FoodColumns struct {
	Id       string // Pirmary ID
	Name     string // Food name
	Price    string // Price
	Quantity string // Quantity
}

// foodColumns holds the columns for table food.
var foodColumns = FoodColumns{
	Id:       "id",
	Name:     "name",
	Price:    "price",
	Quantity: "quantity",
}

// NewFoodDao creates and returns a new dao object for table data access.
func NewFoodDao() *FoodDao {
	return &FoodDao{
		group:   "test2",
		table:   "food",
		columns: foodColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current dao.
func (dao *FoodDao) DB() *gorm.DB { return fmysql.DB(dao.group) }

// Table returns the table name of current dao.
func (dao *FoodDao) Table() string { return dao.table }

// Columns returns all column names of current dao.
func (dao *FoodDao) Columns() FoodColumns { return dao.columns }

// Group returns the configuration group name of database of current dao.
func (dao *FoodDao) Group() string { return dao.group }

// Session creates and returns the session for current dao.
func (dao *FoodDao) Session() *gorm.DB { return dao.DB().Table(dao.table) }

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not commit or rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FoodDao) Transaction(f func(tx *gorm.DB) error) error { return dao.DB().Transaction(f) }
