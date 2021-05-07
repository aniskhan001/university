package repo

import (
	"university/model"
)

type DetailPresenter struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Department  string `json:"department"`
	Designation string `json:"designation"`
}

// Making more readable response object for the users to consume
func ToPresenter(data *model.Teacher) *Presenter {
	return &Presenter{
		ID:          data.ID,
		Name:        data.Name,
		Department:  data.Department,
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
