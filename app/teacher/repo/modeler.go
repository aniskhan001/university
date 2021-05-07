package repo

import (
	"university/model"

	"gorm.io/gorm"
)

type Presenter struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Department  uint   `json:"department"`
	Designation string `json:"designation"`
}

// converting payload to proper model object
func ToModel(data *Presenter) *model.Teacher {
	return &model.Teacher{
		// let DB decide the ID, resetting ID to 0 if provided by client
		Model:       gorm.Model{ID: 0},
		Name:        data.Name,
		Department:  data.Department,
		Designation: data.Designation,
	}
}

func ToModelList(data []Presenter) []model.Teacher {
	res := []model.Teacher{}
	for _, d := range data {
		res = append(res, *ToModel(&d))
	}
	return res
}
