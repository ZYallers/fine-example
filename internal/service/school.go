// ================================================================================
// Code generated and maintained by Fine CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// Created at 2024/10/30 17:47:39.618.
// ================================================================================

package service

import (
	"github.com/ZYallers/fine-example/api/school/student"
	"github.com/ZYallers/fine-example/internal/model"
)

type (
	ISchoolStudent interface {
		Add(req *student.AddReq) (err error)
		Info(id uint) (res *model.StudentInfo)
	}
)

var (
	localSchoolStudent ISchoolStudent
)

func SchoolStudent() ISchoolStudent {
	if localSchoolStudent == nil {
		panic("implement not found for interface ISchoolStudent, forgot register?")
	}
	return localSchoolStudent
}

func RegisterSchoolStudent(i ISchoolStudent) {
	localSchoolStudent = i
}
