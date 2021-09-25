package model

import "gorm.io/gorm"

type Club struct {
	gorm.Model
	Name         string
	DepartmentID uint
	PresidentID  uint
	SecretaryID  uint

	Department Department `gorm:"foreignkey:DepartmentID"`
	President  Student    `gorm:"foreignkey:PresidentID"`
	Secretary  Student    `gorm:"foreignkey:SecretaryID"`
}
