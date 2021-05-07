package repo

import (
	"university/model"

	"gorm.io/gorm"
)

type repo struct {
	*gorm.DB
}

type Repo interface {
	Get(uint) (*model.Club, error)
	GetAll() ([]model.Club, error)
	Insert(*model.Club) (*model.Club, error)
	Edit(uint, *model.Club) (*model.Club, error)
}

// Init will create an object that represent the Repo interface
func Init(DB *gorm.DB) Repo {
	return &repo{
		DB: DB,
	}
}

func (db *repo) Get(id uint) (*model.Club, error) {
	res := model.Club{}

	if op := db.Select("id", "name").Where("id = ?", id).First(&res); op.Error != nil {
		return nil, op.Error
	}

	return &res, nil
}

func (db *repo) GetAll() ([]model.Club, error) {
	res := []model.Club{}

	if op := db.Select("id", "name").Find(&res); op.Error != nil {
		return nil, op.Error
	}

	return res, nil
}

func (db *repo) Insert(data *model.Club) (*model.Club, error) {
	if op := db.Create(data); op.Error != nil {
		return nil, op.Error
	}

	return data, nil
}

func (db *repo) Edit(id uint, data *model.Club) (*model.Club, error) {
	op := db.Model(model.Club{}).Where("id = ?", id).Updates(*data)
	if op.Error != nil {
		return nil, op.Error
	}

	return db.Get(id)
}
