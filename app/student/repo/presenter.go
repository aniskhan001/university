package repo

import (
	deptModel "university/app/department/model"
	"university/app/student/model"
)

type StudentDetailsResp struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Clubs      string `json:"clubs"`
}

// Making more readable response object for the users to consume
func ToStudentResponse(student *model.Student) *StudentResp {
	return &StudentResp{
		ID:         student.ID,
		Name:       student.Name,
		Department: student.Department,
		Clubs:      student.Clubs,
	}
}

func ToStudentsResponse(students []model.Student) []StudentResp {
	res := []StudentResp{}
	for _, student := range students {
		res = append(res, *ToStudentResponse(&student))
	}
	return res
}

func StudentDetailsResponse(student *model.Student, dept *deptModel.Department) *StudentDetailsResp {
	return &StudentDetailsResp{
		ID:         student.ID,
		Name:       student.Name,
		Department: dept.Name,
		Clubs:      student.Clubs,
	}
}
