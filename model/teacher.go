package model

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name         string
	DepartmentID uint
	Designation  string

	Department Department `gorm:"foreignkey:DepartmentID"`
}
