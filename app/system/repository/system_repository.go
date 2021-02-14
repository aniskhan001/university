package repository

import (
	"kabikha/infrastructure/db"
	"time"

	"github.com/jinzhu/gorm"
)

type pgSystemRepository struct {
	*gorm.DB
}

// NewPgSystemRepository will create an object that represent the article.Repository interface
func NewPgSystemRepository(DB *gorm.DB) SystemRepository {
	return &pgSystemRepository{
		DB: db.Get().DB,
	}
}

type SystemRepository interface {
	DBCheck() (bool, error)
	CurrentTime() int64
}

func (db *pgSystemRepository) DBCheck() (bool, error) {
	if err := db.DB.DB().Ping(); err != nil {
		return false, err
	}
	return true, nil
}

func (db *pgSystemRepository) CurrentTime() int64 {
	return time.Now().Unix()
}
