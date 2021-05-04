package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name       string `json:"name"`
	Department uint   `json:"department"`
	Clubs      string `json:"clubs"`
}
