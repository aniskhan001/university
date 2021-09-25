package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name         string
	Roll         uint
	DepartmentID uint

	Department Department `gorm:"foreignkey:DepartmentID"`
}
