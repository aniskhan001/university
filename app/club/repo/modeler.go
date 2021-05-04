package repo

import (
	"university/model"

	"gorm.io/gorm"
)

type Response struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Department uint   `json:"department"`
	President  uint   `json:"president"`
	Secretary  uint   `json:"secretary"`
}

// converting payload to proper model object
func ToModel(data *Response) *model.Club {
	return &model.Club{
		// let DB decide the ID, resetting ID to 0 if provided by client
		Model:      gorm.Model{ID: 0},
		Name:       data.Name,
		Department: data.Department,
		President:  data.President,
		Secretary:  data.Secretary,
	}
}
