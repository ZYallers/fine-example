// =================================================================================
// Code generated and maintained by Fine CLI tool. DO NOT EDIT.
// Created at 2024/10/30 18:36:15.435.
// =================================================================================

package school

import (
	"github.com/gin-gonic/gin"

	"github.com/ZYallers/fine-example/api/school/student"
)

type ISchoolStudent interface {
	Add(ctx *gin.Context, req *student.AddReq) (res *student.AddRes, err error)
	Info(ctx *gin.Context, req *student.InfoReq) (res *student.InfoRes, err error)
	QueryScore(ctx *gin.Context, req *student.QueryScoreReq) (res *student.QueryScoreRes, err error)
}
