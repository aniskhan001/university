package usecase

import "university/app/department/model"

type DeptResp struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// Making more readable response object for the users to consume
func toDeptResp(dept *model.Dept) *DeptResp {
	return &DeptResp{
		ID:   dept.ID,
		Name: dept.Name,
	}
}

func toDeptResps(depts []model.Dept) []DeptResp {
	res := []DeptResp{}
	for _, dept := range depts {
		res = append(res, *toDeptResp(&dept))
	}
	return res
}
