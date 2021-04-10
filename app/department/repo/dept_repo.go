package repo

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
	Insert(*model.Dept) (*model.Dept, error)
	InsertMany([]model.Dept) ([]model.Dept, error)
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

	if res := db.Select("id", "name").Where("id = ?", id).First(&dept); res.Error != nil {
		return nil, res.Error
	}

	return &dept, nil
}

func (db *deptRepository) GetAll() ([]model.Dept, error) {
	var depts []model.Dept

	if res := db.Select("id", "name").Find(&depts); res.Error != nil {
		return nil, res.Error
	}

	return depts, nil
}

func (db *deptRepository) Insert(dept *model.Dept) (*model.Dept, error) {
	if res := db.Create(dept); res.Error != nil {
		return nil, res.Error
	}

	return dept, nil
}

func (db *deptRepository) InsertMany(depts []model.Dept) ([]model.Dept, error) {
	if res := db.Create(depts); res.Error != nil {
		return nil, res.Error
	}

	return depts, nil
}
