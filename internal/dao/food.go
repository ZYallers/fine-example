// =================================================================================
// This is auto-generated by Fine CLI tool only once. Fill this file as you wish.
// Created at 2024/10/30 17:57:57.865.
// =================================================================================

package dao

import (
	"github.com/ZYallers/fine-example/internal/dao/internal"
)

// internalFoodDao is internal type for wrapping internal dao implements.
type internalFoodDao = *internal.FoodDao

// foodDao is the data access object for table food.
// You can define custom methods on it to extend its functionality as you wish.
type foodDao struct {
	internalFoodDao
}

var (
	// Food is globally public accessible object for table food operations.
	Food = foodDao{
		internal.NewFoodDao(),
	}
)
