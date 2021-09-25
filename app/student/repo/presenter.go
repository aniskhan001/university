package repo

import (
	"university/model"

	"gorm.io/gorm"
)

type Presenter struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	DepartmentID uint   `json:"department_id"`
	Department   string `json:"department"`
}

func ToPresenter(data *model.Student) *Presenter {
	return &Presenter{
		ID:           data.ID,
		Name:         data.Name,
		DepartmentID: data.DepartmentID,
	}
}

func ToPresenterList(data []model.Student) []Presenter {
	res := []Presenter{}
	for _, d := range data {
		res = append(res, *ToPresenter(&d))
	}
	return res
}

func ToModel(data *Presenter) *model.Student {
	return &model.Student{
		// let DB decide the ID, resetting ID to 0 if provided by client
		Model:        gorm.Model{ID: 0},
		Name:         data.Name,
		DepartmentID: data.DepartmentID,
	}
}

func ToModelList(data []Presenter) []model.Student {
	res := []model.Student{}
	for _, d := range data {
		res = append(res, *ToModel(&d))
	}
	return res
}
