package repo

import (
	"gorm.io/gorm"
)

type repo struct {
	*gorm.DB
}

// Init will create an object that represent the Repo interface
func Init(DB *gorm.DB) Repo {
	return &repo{
		DB: DB,
	}
}

type Repo interface {
	DBCheck() (bool, error)
}

func (db *repo) DBCheck() (bool, error) {
	dbInstance, err := db.DB.DB()
	if err != nil {
		return false, err
	}
	if err = dbInstance.Ping(); err != nil {
		return false, err
	}
	return true, nil
}
