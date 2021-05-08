package usecase

import "university/app/system/repo"

type Usecase interface {
	GetHealth() (*HealthResp, error)
}

type usecase struct {
	repo repo.Repo
}

func Init(repo repo.Repo) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) GetHealth() (*HealthResp, error) {
	resp := HealthResp{}

	// check db
	db_online, err := uc.repo.DBCheck()
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
