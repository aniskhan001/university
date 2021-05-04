package repo

import (
	"university/model"

	"gorm.io/gorm"
)

type DeptResp struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// converting payload to proper model object
func ToDeptModel(dr *DeptResp) *model.Department {
	return &model.Department{
		// let DB decide the ID, resetting ID to 0 if provided by client
		Model: gorm.Model{ID: 0},
		Name:  dr.Name,
	}
}

func ToDeptsModel(drs []DeptResp) []model.Department {
	res := []model.Department{}
	for _, dr := range drs {
		res = append(res, *ToDeptModel(&dr))
	}
	return res
}
