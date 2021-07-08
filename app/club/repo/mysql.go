package repo

import (
	"university/model"

	"gorm.io/gorm"
)

type mysqlRepo struct {
	*gorm.DB
}

// InitMySQL will create an object that represent the Repo interface
func InitMySQL(DB *gorm.DB) Repo {
	return &mysqlRepo{
		DB: DB,
	}
}

func (db *mysqlRepo) Get(id uint) (*model.Club, error) {
	res := model.Club{}

	if op := db.Select("id", "name").Where("id = ?", id).First(&res); op.Error != nil {
		return nil, op.Error
	}

	return &res, nil
}

func (db *mysqlRepo) GetAll() ([]model.Club, error) {
	res := []model.Club{}

	if op := db.Select("id", "name").Find(&res); op.Error != nil {
		return nil, op.Error
	}

	return res, nil
}

func (db *mysqlRepo) Insert(data *model.Club) (*model.Club, error) {
	if op := db.Create(data); op.Error != nil {
		return nil, op.Error
	}

	return data, nil
}

func (db *mysqlRepo) Edit(id uint, data *model.Club) (*model.Club, error) {
	op := db.Model(model.Club{}).Where("id = ?", id).Updates(*data)
	if op.Error != nil {
		return nil, op.Error
	}

	return db.Get(id)
}
