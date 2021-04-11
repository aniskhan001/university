package repo

import "university/app/teacher/model"

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
