// =================================================================================
// This is auto-generated by Fine CLI tool only once. Fill this file as you wish.
// Created at 2024/10/30 17:40:53.201.
// =================================================================================

package student

import (
	"github.com/ZYallers/fine-example/api/school"
)

type cStudent struct{}

func New() school.ISchoolStudent {
	return &cStudent{}
}
