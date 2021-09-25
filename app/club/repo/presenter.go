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
	PresidentID  uint   `json:"president_id"`
	President    string `json:"president"`
	SecretaryID  uint   `json:"secretary_id"`
	Secretary    string `json:"secretary"`
}

func ToPresenter(data *model.Club) *Presenter {
	return &Presenter{
		ID:           data.ID,
		Name:         data.Name,
		DepartmentID: data.DepartmentID,
		PresidentID:  data.PresidentID,
		SecretaryID:  data.SecretaryID,
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
		Model:        gorm.Model{ID: 0},
		Name:         data.Name,
		DepartmentID: data.DepartmentID,
		PresidentID:  data.PresidentID,
		SecretaryID:  data.SecretaryID,
	}
}
