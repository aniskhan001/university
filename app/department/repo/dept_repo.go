package repo

import (
	"university/infrastructure/db"
	"university/model"

	"gorm.io/gorm"
)

type deptRepository struct {
	*gorm.DB
}

type DeptRepository interface {
	Get(uint) (*model.Department, error)
	GetAll() ([]model.Department, error)
	Insert(*model.Department) (*model.Department, error)
	InsertMany([]model.Department) ([]model.Department, error)
	Edit(uint, *model.Department) (*model.Department, error)
}

// NewDeptRepository will create an object that represent the DeptRepository interface
func NewDeptRepository(DB *gorm.DB) DeptRepository {
	return &deptRepository{
		DB: db.Get().DB,
	}
}

func (db *deptRepository) Get(id uint) (*model.Department, error) {
	var dept model.Department

	if res := db.Select("id", "name").Where("id = ?", id).First(&dept); res.Error != nil {
		return nil, res.Error
	}

	return &dept, nil
}

func (db *deptRepository) GetAll() ([]model.Department, error) {
	var depts []model.Department

	if res := db.Select("id", "name").Find(&depts); res.Error != nil {
		return nil, res.Error
	}

	return depts, nil
}

func (db *deptRepository) Insert(dept *model.Department) (*model.Department, error) {
	if res := db.Create(dept); res.Error != nil {
		return nil, res.Error
	}

	return dept, nil
}

func (db *deptRepository) InsertMany(depts []model.Department) ([]model.Department, error) {
	if res := db.Create(depts); res.Error != nil {
		return nil, res.Error
	}

	return depts, nil
}

func (db *deptRepository) Edit(id uint, dept *model.Department) (*model.Department, error) {
	res := db.Model(model.Department{}).Where("id = ?", id).Updates(*dept)
	if res.Error != nil {
		return nil, res.Error
	}

	return db.Get(id)
}
