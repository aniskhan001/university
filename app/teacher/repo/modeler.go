package repo

import (
	"university/app/teacher/model"

	"gorm.io/gorm"
)

type TeacherResp struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Department  uint   `json:"department"`
	Designation string `json:"designation"`
}

// converting payload to proper model object
func ToTeacherModel(tr *TeacherResp) *model.Teacher {
	return &model.Teacher{
		// let DB decide the ID, resetting ID to 0 if provided by client
		Model:       gorm.Model{ID: 0},
		Name:        tr.Name,
		Department:  tr.Department,
		Designation: tr.Designation,
	}
}

func ToTeachersModel(trs []TeacherResp) []model.Teacher {
	res := []model.Teacher{}
	for _, dr := range trs {
		res = append(res, *ToTeacherModel(&dr))
	}
	return res
}
