// =================================================================================
// This is auto-generated by Fine CLI tool only once. Fill this file as you wish.
// Created at 2024/10/30 17:40:53.197.
// =================================================================================

package device

import (
	"github.com/ZYallers/fine-example/api/product"
)

type cDevice struct{}

func New() product.IProductDevice {
	return &cDevice{}
}