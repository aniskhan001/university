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

type DetailPresenter struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Department  string `json:"department"`
	Designation string `json:"designation"`
}

func ToPresenter(data *model.Teacher) *Presenter {
	return &Presenter{
		ID:          data.ID,
		Name:        data.Name,
		Department:  data.DepartmentID,
		Designation: data.Designation,
	}
}

func ToPresenterList(data []model.Teacher) []Presenter {
	res := []Presenter{}
	for _, d := range data {
		res = append(res, *ToPresenter(&d))
	}
	return res
}

func ToDetailPresenter(data *model.Teacher, deptData *model.Department) *DetailPresenter {
	return &DetailPresenter{
		ID:          data.ID,
		Name:        data.Name,
		Department:  deptData.Name,
		Designation: data.Designation,
	}
}

func ToModel(data *Presenter) *model.Teacher {
	return &model.Teacher{
		// let DB decide the ID, resetting ID to 0 if provided by client
		Model:        gorm.Model{ID: 0},
		Name:         data.Name,
		DepartmentID: data.Department,
		Designation:  data.Designation,
	}
}

func ToModelList(data []Presenter) []model.Teacher {
	res := []model.Teacher{}
	for _, d := range data {
		res = append(res, *ToModel(&d))
	}
	return res
}
