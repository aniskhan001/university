package repo

import (
	"university/model"

	"gorm.io/gorm"
)

type repo struct {
	*gorm.DB
}

type Repo interface {
	Get(uint) (*model.Student, error)
	Insert(*model.Student) (*model.Student, error)
	InsertMany([]model.Student) ([]model.Student, error)
	GetAll() ([]model.Student, error)
	GetAllFromDept(uint) ([]model.Student, error)
	Edit(uint, *model.Student) (*model.Student, error)
}

// Init will create an object that represent the Repo interface
func Init(DB *gorm.DB) Repo {
	return &repo{
		DB: DB,
	}
}

var fields = []string{"id", "name", "roll", "department_id"}

func (db *repo) Get(id uint) (*model.Student, error) {
	res := model.Student{}

	if op := db.Select(fields).Where("id = ?", id).First(&res); op.Error != nil {
		return nil, op.Error
	}

	return &res, nil
}

func (db *repo) GetAll() ([]model.Student, error) {
	res := []model.Student{}

	if op := db.Select(fields).Find(&res); op.Error != nil {
		return nil, op.Error
	}

	return res, nil
}

func (db *repo) GetAllFromDept(deptID uint) ([]model.Student, error) {
	res := []model.Student{}

	if op := db.Select(fields).Where("department = ?", deptID).Find(&res); op.Error != nil {
		return nil, op.Error
	}

	return res, nil
}

func (db *repo) Insert(data *model.Student) (*model.Student, error) {
	if op := db.Create(data); op.Error != nil {
		return nil, op.Error
	}

	return data, nil
}

func (db *repo) InsertMany(data []model.Student) ([]model.Student, error) {
	if op := db.Create(data); op.Error != nil {
		return nil, op.Error
	}

	return data, nil
}

func (db *repo) Edit(id uint, data *model.Student) (*model.Student, error) {
	if op := db.Model(model.Student{}).Where("id = ?", id).Updates(*data); op.Error != nil {
		return nil, op.Error
	}

	return db.Get(id)
}
