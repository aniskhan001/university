package repo

import (
	"gorm.io/gorm"
)

type systemRepository struct {
	*gorm.DB
}

// NewSystemRepository will create an object that represent the article.Repository interface
func NewSystemRepository(DB *gorm.DB) SystemRepository {
	return &systemRepository{
		DB: DB,
	}
}

type SystemRepository interface {
	DBCheck() (bool, error)
}

func (db *systemRepository) DBCheck() (bool, error) {
	dbInstance, err := db.DB.DB()
	if err != nil {
		return false, err
	}
	if err = dbInstance.Ping(); err != nil {
		return false, err
	}
	return true, nil
}
