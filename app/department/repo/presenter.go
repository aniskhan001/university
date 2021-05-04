package repo

import "university/model"

// Making more readable response object for the users to consume
func ToDeptResponse(dept *model.Department) *DeptResp {
	return &DeptResp{
		ID:   dept.ID,
		Name: dept.Name,
	}
}

func ToDeptsResponse(depts []model.Department) []DeptResp {
	res := []DeptResp{}
	for _, dept := range depts {
		res = append(res, *ToDeptResponse(&dept))
	}
	return res
}
