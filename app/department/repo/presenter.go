package repo

import (
	"university/model"

	"gorm.io/gorm"
)

type Presenter struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ToPresenter(data *model.Department) *Presenter {
	return &Presenter{
		ID:   data.ID,
		Name: data.Name,
	}
}

func ToPresenterList(data []model.Department) []Presenter {
	res := []Presenter{}
	for _, d := range data {
		res = append(res, *ToPresenter(&d))
	}
	return res
}

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
