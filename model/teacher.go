package model

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name        string `json:"name"`
	Department  uint   `json:"department"`
	Designation string `json:"designation"`
}
