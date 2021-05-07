package repo

import (
	"university/model"
)

type DetailPresenter struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Clubs      string `json:"clubs"`
}

// Making more readable response object for the users to consume
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
