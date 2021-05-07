package repo

import (
	"university/model"

	"gorm.io/gorm"
)

type Presenter struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Department uint   `json:"department"`
	President  uint   `json:"president"`
	Secretary  uint   `json:"secretary"`
}

func ToPresenter(data *model.Club) *Presenter {
	return &Presenter{
		ID:         data.ID,
		Name:       data.Name,
		Department: data.Department,
		President:  data.President,
		Secretary:  data.Secretary,
	}
}

func ToPresenterList(data []model.Club) []Presenter {
	res := []Presenter{}
	for _, d := range data {
		res = append(res, *ToPresenter(&d))
	}
	return res
}

func ToModel(data *Presenter) *model.Club {
	return &model.Club{
		// let DB decide the ID, resetting ID to 0 if provided by client
		Model:      gorm.Model{ID: 0},
		Name:       data.Name,
		Department: data.Department,
		President:  data.President,
		Secretary:  data.Secretary,
	}
}
