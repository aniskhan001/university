package repo

import (
	"university/app/student/model"

	"gorm.io/gorm"
)

type StudentResp struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Department uint   `json:"department"`
	Clubs      string `json:"clubs"`
}

// converting payload to proper model object
func ToStudentModel(sr *StudentResp) *model.Student {
	return &model.Student{
		// let DB decide the ID, resetting ID to 0 if provided by client
		Model:      gorm.Model{ID: 0},
		Name:       sr.Name,
		Department: sr.Department,
		Clubs:      sr.Clubs,
	}
}

func ToStudentsModel(srs []StudentResp) []model.Student {
	res := []model.Student{}
	for _, dr := range srs {
		res = append(res, *ToStudentModel(&dr))
	}
	return res
}
