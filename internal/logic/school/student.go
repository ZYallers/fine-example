package school

import (
	"github.com/ZYallers/fine-example/api/school/student"
	"github.com/ZYallers/fine-example/internal/model"
	"github.com/ZYallers/fine-example/internal/service"
)

type sStudent struct {
}

func init() {
	service.RegisterSchoolStudent(NewStudent())
}

func NewStudent() *sStudent {
	s := &sStudent{}
	return s
}

// Add New Student add
func (s *sStudent) Add(req *student.AddReq) (err error) {
	return
}

// Info Student info
func (s *sStudent) Info(id uint) (res *model.StudentInfo) {
	return
}
