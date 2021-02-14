package repository

import (
	"kabikha/infrastructure/db"

	"github.com/jinzhu/gorm"
)

type systemRepository struct {
	*gorm.DB
}

// NewSystemRepository will create an object that represent the article.Repository interface
func NewSystemRepository(DB *gorm.DB) SystemRepository {
	return &systemRepository{
		DB: db.Get().DB,
	}
}

type SystemRepository interface {
	DBCheck() (bool, error)
}

func (db *systemRepository) DBCheck() (bool, error) {
	if err := db.DB.DB().Ping(); err != nil {
		return false, err
	}
	return true, nil
}
