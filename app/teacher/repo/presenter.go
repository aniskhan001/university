package repo

import (
	deptModel "university/app/department/model"
	"university/app/teacher/model"
)

type TeacherDetailsResp struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Department  string `json:"department"`
	Designation string `json:"designation"`
}

// Making more readable response object for the users to consume
func ToTeacherResponse(teacher *model.Teacher) *TeacherResp {
	return &TeacherResp{
		ID:          teacher.ID,
		Name:        teacher.Name,
		Department:  teacher.Department,
		Designation: teacher.Designation,
	}
}

func ToTeachersResponse(teachers []model.Teacher) []TeacherResp {
	res := []TeacherResp{}
	for _, teacher := range teachers {
		res = append(res, *ToTeacherResponse(&teacher))
	}
	return res
}

func TeacherDetailsResponse(teacher *model.Teacher, dept *deptModel.Department) *TeacherDetailsResp {
	return &TeacherDetailsResp{
		ID:          teacher.ID,
		Name:        teacher.Name,
		Department:  dept.Name,
		Designation: teacher.Designation,
	}
}
