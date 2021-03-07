package model

import "gorm.io/gorm"

type Dept struct {
	gorm.Model
	Name string `json:"name"`
}
