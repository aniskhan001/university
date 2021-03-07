package repository

import (
	"university/app/department/model"
	"university/infrastructure/db"

	"gorm.io/gorm"
)

type deptRepository struct {
	*gorm.DB
}

type DeptRepository interface {
	Get(uint) (*model.Dept, error)
	GetAll() ([]model.Dept, error)
}

// NewDeptRepository will create an object that represent the DeptRepository interface
func NewDeptRepository(DB *gorm.DB) DeptRepository {
	return &deptRepository{
		DB: db.Get().DB,
	}
}

func (db *deptRepository) Get(id uint) (*model.Dept, error) {
	var dept model.Dept

	res := db.Select("id", "name").Where("id = ?", id).First(&dept)
	if res.Error != nil {
		return nil, res.Error
	}

	return &dept, nil
}

func (db *deptRepository) GetAll() ([]model.Dept, error) {
	var depts []model.Dept
	res := db.Select("id", "name").Find(&depts)
	if res.Error != nil {
		return nil, res.Error
	}
	return depts, nil
}
