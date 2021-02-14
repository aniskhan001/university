package usecase

import (
	"kabikha/app/system/repository"
)

type SystemUsecase interface {
	GetHealth() (*HealthResp, error)
}

type systemUsecase struct {
	repo repository.SystemRepository
}

func NewSystemUsecase(repo repository.SystemRepository) SystemUsecase {
	return &systemUsecase{
		repo: repo,
	}
}

func (u *systemUsecase) GetHealth() (*HealthResp, error) {
	resp := HealthResp{}

	// check db
	db_online, err := u.repo.DBCheck()
	resp.DBOnline = db_online
	if err != nil {
		return &resp, err
	}

	return &resp, nil
}

type HealthResp struct {
	DBOnline    bool `json:"db_online"`
	CacheOnline bool `json:"cache_online"`
}
