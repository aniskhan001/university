package repo

import (
	"university/model"

	"gorm.io/gorm"
)

type repo struct {
	*gorm.DB
}

type Repo interface {
	Get(uint) (*model.Department, error)
	GetAll() ([]model.Department, error)
	Insert(*model.Department) (*model.Department, error)
	InsertMany([]model.Department) ([]model.Department, error)
	Edit(uint, *model.Department) (*model.Department, error)
}

// Init will create an object that represent the Repo interface
func Init(DB *gorm.DB) Repo {
	return &repo{
		DB: DB,
	}
}

func (db *repo) Get(id uint) (*model.Department, error) {
	var res model.Department

	if op := db.Select("id", "name").Where("id = ?", id).First(&res); op.Error != nil {
		return nil, op.Error
	}

	return &res, nil
}

func (db *repo) GetAll() ([]model.Department, error) {
	var res []model.Department

	if op := db.Select("id", "name").Find(&res); op.Error != nil {
		return nil, op.Error
	}

	return res, nil
}

func (db *repo) Insert(data *model.Department) (*model.Department, error) {
	if op := db.Create(data); op.Error != nil {
		return nil, op.Error
	}

	return data, nil
}

func (db *repo) InsertMany(data []model.Department) ([]model.Department, error) {
	if op := db.Create(data); op.Error != nil {
		return nil, op.Error
	}

	return data, nil
}

func (db *repo) Edit(id uint, data *model.Department) (*model.Department, error) {
	if op := db.Model(model.Department{}).Where("id = ?", id).Updates(*data); op.Error != nil {
		return nil, op.Error
	}

	return db.Get(id)
}
