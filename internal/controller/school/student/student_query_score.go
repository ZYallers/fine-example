// =================================================================================
// This is auto-generated by Fine CLI tool only once. Fill this file as you wish.
// Created at 2024/10/30 18:36:15.435.
// =================================================================================

package student

import (
	"github.com/gin-gonic/gin"

	"github.com/ZYallers/fine-example/api/school/student"
)

//	@Summary	school/student/query/score
//	@Tags		school
//	@Accept		json
//	@Produce	json
//	@Router		/school/student/query/score [post]
//	@Param		params	body		student.QueryScoreReq	true	"..."
//	@Success	200		{object}	f.JsonResult{data=student.QueryScoreRes}
func (c *cStudent) QueryScore(ctx *gin.Context, req *student.QueryScoreReq) (res *student.QueryScoreRes, err error) {
	return
}
