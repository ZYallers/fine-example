package student

import "github.com/ZYallers/fine/frame/f"

type QueryScoreReq struct {
	f.Meta    `path:"student/student/query/score" method:"post"`
	SubjectId uint `form:"subject_id" validate:"required,min=1" json:"subject_id"` // subject id
}

type QueryScoreRes struct{}
