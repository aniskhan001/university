package repo

import (
	"university/model"

	"gorm.io/gorm"
)

type Presenter struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Department uint   `json:"department"`
	Clubs      string `json:"clubs"`
}

// converting payload to proper model object
func ToModel(data *Presenter) *model.Student {
	return &model.Student{
		// let DB decide the ID, resetting ID to 0 if provided by client
		Model:      gorm.Model{ID: 0},
		Name:       data.Name,
		Department: data.Department,
		Clubs:      data.Clubs,
	}
}

func ToModelList(data []Presenter) []model.Student {
	res := []model.Student{}
	for _, d := range data {
		res = append(res, *ToModel(&d))
	}
	return res
}
