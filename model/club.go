package model

import "gorm.io/gorm"

type Club struct {
	gorm.Model
	Name       string `json:"name"`
	Department uint   `json:"department"`
	President  uint   `json:"president"`
	Secretary  uint   `json:"secretary"`
}
