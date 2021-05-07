package repo

import (
	"university/model"

	"gorm.io/gorm"
)

type Presenter struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// converting payload to proper model object
func ToModel(dr *Presenter) *model.Department {
	return &model.Department{
		// let DB decide the ID, resetting ID to 0 if provided by client
		Model: gorm.Model{ID: 0},
		Name:  dr.Name,
	}
}

func ToModelList(drs []Presenter) []model.Department {
	res := []model.Department{}
	for _, dr := range drs {
		res = append(res, *ToModel(&dr))
	}
	return res
}
