package model

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name string `json:"name"`
}
