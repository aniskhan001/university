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

type DetailPresenter struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Clubs      string `json:"clubs"`
}

func ToPresenter(data *model.Student) *Presenter {
	return &Presenter{
		ID:         data.ID,
		Name:       data.Name,
		Department: data.Department,
		Clubs:      data.Clubs,
	}
}

func ToPresenterList(data []model.Student) []Presenter {
	res := []Presenter{}
	for _, d := range data {
		res = append(res, *ToPresenter(&d))
	}
	return res
}

func ToDetailPresenter(data *model.Student, deptData *model.Department) *DetailPresenter {
	return &DetailPresenter{
		ID:         data.ID,
		Name:       data.Name,
		Department: deptData.Name,
		Clubs:      data.Clubs,
	}
}

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
