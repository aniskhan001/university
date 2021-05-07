package repo

import (
	"university/model"

	"gorm.io/gorm"
)

type repo struct {
	*gorm.DB
}

type Repo interface {
	Get(uint) (*model.Teacher, error)
	Insert(*model.Teacher) (*model.Teacher, error)
	InsertMany([]model.Teacher) ([]model.Teacher, error)
	GetAll() ([]model.Teacher, error)
	GetAllFromDept(uint) ([]model.Teacher, error)
	Edit(uint, *model.Teacher) (*model.Teacher, error)
}

// Init will create an object that represent the Repo interface
func Init(DB *gorm.DB) Repo {
	return &repo{
		DB: DB,
	}
}

func (db *repo) Get(id uint) (*model.Teacher, error) {
	res := model.Teacher{}

	if op := db.Select("id", "name", "department").Where("id = ?", id).First(&res); op.Error != nil {
		return nil, op.Error
	}

	return &res, nil
}

func (db *repo) GetAll() ([]model.Teacher, error) {
	res := []model.Teacher{}

	if op := db.Select("id", "name").Find(&res); op.Error != nil {
		return nil, op.Error
	}

	return res, nil
}

func (db *repo) GetAllFromDept(deptID uint) ([]model.Teacher, error) {
	res := []model.Teacher{}

	if op := db.Select("id", "name").Where("department = ?", deptID).Find(&res); op.Error != nil {
		return nil, op.Error
	}

	return res, nil
}

func (db *repo) Insert(data *model.Teacher) (*model.Teacher, error) {
	if op := db.Create(data); op.Error != nil {
		return nil, op.Error
	}

	return data, nil
}

func (db *repo) InsertMany(data []model.Teacher) ([]model.Teacher, error) {
	if op := db.Create(data); op.Error != nil {
		return nil, op.Error
	}

	return data, nil
}

func (db *repo) Edit(id uint, teacher *model.Teacher) (*model.Teacher, error) {
	if op := db.Model(model.Teacher{}).Where("id = ?", id).Updates(*teacher); op.Error != nil {
		return nil, op.Error
	}

	return db.Get(id)
}
